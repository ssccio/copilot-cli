// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package multi_pipeline_test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aws/copilot-cli/e2e/internal/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("pipeline flow", func() {
	Context("set up CodeCommit repository", func() {
		It("creates the codecommit repository", func() {
			url, err := aws.CreateCodeCommitRepo(repoName)
			Expect(err).NotTo(HaveOccurred())
			repoURL = url
		})

		It("clones the repository", func() {
			endpoint := strings.TrimPrefix(repoURL, "https://")
			url := fmt.Sprintf("https://%s:%s@%s", url.PathEscape(codeCommitCreds.UserName), url.PathEscape(codeCommitCreds.Password), endpoint)

			Eventually(func() error {
				cmd := exec.Command("git", "clone", url)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				return cmd.Run()
			}, "60s", "5s").ShouldNot(HaveOccurred())
		})

		It("copies source code to the git repository", func() {
			cmd := exec.Command("cp", "-r", "frontend", repoName)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})

		It("sets git config", func() {
			cmd := exec.Command("git", "config", "user.email", "e2etest@amazon.com")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())

			cmd = exec.Command("git", "config", "user.name", "e2etest")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
	})

	Context("create a new app", func() {
		It("app init succeeds", func() {
			_, err := copilot.AppInit(&client.AppInitRequest{
				AppName: appName,
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("app init creates an copilot directory and workspace file", func() {
			Expect(filepath.Join(repoName, "copilot")).Should(BeADirectory())
			Expect(filepath.Join(repoName, "copilot", ".workspace")).Should(BeAnExistingFile())
		})
		It("app ls includes new app", func() {
			Eventually(copilot.AppList, "30s", "5s").Should(ContainSubstring(appName))
		})
		It("app show includes app name", func() {
			appShowOutput, err := copilot.AppShow(appName)
			Expect(err).NotTo(HaveOccurred())
			Expect(appShowOutput.Name).To(Equal(appName))
			Expect(appShowOutput.URI).To(BeEmpty())
		})
	})

	Context("when adding a new environment", func() {
		It("test env init should succeed", func() {
			_, err := copilot.EnvInit(&client.EnvInitRequest{
				AppName: appName,
				EnvName: "test",
				Profile: "e2etestenv",
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("prod env init should succeed", func() {
			_, err := copilot.EnvInit(&client.EnvInitRequest{
				AppName: appName,
				EnvName: "prod",
				Profile: "e2eprodenv",
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("env ls should list both envs", func() {
			out, err := copilot.EnvList(appName)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(out.Envs)).To(Equal(2))
			envs := map[string]client.EnvDescription{}
			for _, env := range out.Envs {
				envs[env.Name] = env
				Expect(env.ExecutionRole).NotTo(BeEmpty())
				Expect(env.ManagerRole).NotTo(BeEmpty())
			}

			Expect(envs["test"]).NotTo(BeNil())
			Expect(envs["prod"]).NotTo(BeNil())

			// Make sure, for the sake of coverage, these are cross account,
			// cross region environments if we're not doing a dryrun.
			if os.Getenv("DRYRUN") != "true" {
				Expect(envs["test"].Region).NotTo(Equal(envs["prod"].Region))
				Expect(envs["test"].Account).NotTo(Equal(envs["prod"].Account))
			}
		})
	})

	Context("when deploying the environments", func() {
		It("test env deploy should succeed", func() {
			_, err := copilot.EnvDeploy(&client.EnvDeployRequest{
				AppName: appName,
				Name:    "test",
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("prod env deploy should succeed", func() {
			_, err := copilot.EnvDeploy(&client.EnvDeployRequest{
				AppName: appName,
				Name:    "prod",
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when creating the frontend service", func() {
		It("should initialize the service", func() {
			_, err := copilot.SvcInit(&client.SvcInitRequest{
				Name:       "frontend",
				SvcType:    "Load Balanced Web Service",
				Dockerfile: "./frontend/Dockerfile",
				SvcPort:    "80",
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should generate a manifest file", func() {
			Expect(filepath.Join(repoName, "copilot", "frontend", "manifest.yml")).Should(BeAnExistingFile())
		})
		It("creates a new addons dir", func() {
			cmd := exec.Command("mkdir", "-p", filepath.Join(repoName, "copilot", "frontend", "addons"))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
		It("copies a template to the addons file", func() {
			cmd := exec.Command("cp", "s3template.yml", filepath.Join(repoName, "copilot", "frontend", "addons", "e2e-pipeline-addon.yml"))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
		It("should list the service", func() {
			out, err := copilot.SvcList(appName)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(out.Services)).To(Equal(1))
			Expect(out.Services[0].Name).To(Equal("frontend"))
		})
	})

	Context("when creating the test pipeline manifest", func() {
		It("creates a 'test' git branch", func() {
			cmd := exec.Command("git", "checkout", "-b", "test")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
		It("should initialize the pipeline", func() {
			_, err := copilot.PipelineInit(client.PipelineInitInput{
				Name:         testPipelineName,
				URL:          repoURL,
				GitBranch:    "test",
				Environments: []string{"test"},
				Type:         "Workloads",
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should generate pipeline artifacts", func() {
			Expect(filepath.Join(repoName, "copilot", "pipelines", testPipelineName, "manifest.yml")).Should(BeAnExistingFile())
			Expect(filepath.Join(repoName, "copilot", "pipelines", testPipelineName, "buildspec.yml")).Should(BeAnExistingFile())
		})
		It("should push repo changes upstream", func() {
			cmd := exec.Command("git", "add", ".")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())

			cmd = exec.Command("git", "commit", "-m", "first commit")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())

			cmd = exec.Command("git", "push", "--set-upstream", "origin", "test")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
	})

	Context("when creating the prod pipeline manifest", func() {
		It("creates a 'prod' git branch", func() {
			cmd := exec.Command("git", "checkout", "-b", "prod")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
		It("should initialize the pipeline", func() {
			_, err := copilot.PipelineInit(client.PipelineInitInput{
				Name:         prodPipelineName,
				URL:          repoURL,
				GitBranch:    "prod",
				Environments: []string{"prod"},
				Type:         "Workloads",
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should generate pipeline artifacts", func() {
			Expect(filepath.Join(repoName, "copilot", "pipelines", prodPipelineName, "manifest.yml")).Should(BeAnExistingFile())
			Expect(filepath.Join(repoName, "copilot", "pipelines", prodPipelineName, "buildspec.yml")).Should(BeAnExistingFile())
		})

		It("should push repo changes upstream", func() {
			cmd := exec.Command("git", "add", ".")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())

			cmd = exec.Command("git", "commit", "-m", "first commit")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())

			cmd = exec.Command("git", "push", "--set-upstream", "origin", "prod")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
	})

	Context("when creating the test pipeline stack", func() {
		It("checks out the test git branch", func() {
			cmd := exec.Command("git", "checkout", "test")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
		It("should start creating the pipeline stack", func() {
			_, err := copilot.PipelineDeploy(client.PipelineDeployInput{
				Name: testPipelineName,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should show test pipeline details once the stack is created", func() {
			type stage struct {
				Name     string
				Category string
			}
			wantedStages := []stage{
				{
					Name:     "Source",
					Category: "Source",
				},
				{
					Name:     "Build",
					Category: "Build",
				},
				{
					Name:     "DeployTo-test",
					Category: "Deploy",
				},
			}

			Eventually(func() error {
				out, err := copilot.PipelineShow(client.PipelineShowInput{
					Name: testPipelineName,
				})
				switch {
				case err != nil:
					return err
				case out.Name == "":
					return fmt.Errorf("pipeline name is empty: %v", out)
				case out.Name != testPipelineName:
					return fmt.Errorf("expected pipeline name %q, got %q", testPipelineName, out.Name)
				case len(out.Stages) != len(wantedStages):
					return fmt.Errorf("pipeline stages do not match: %v", out.Stages)
				}

				for idx, actualStage := range out.Stages {
					if wantedStages[idx].Name != actualStage.Name {
						return fmt.Errorf("stage name %s at index %d does not match", actualStage.Name, idx)
					}
					if wantedStages[idx].Category != actualStage.Category {
						return fmt.Errorf("stage category %s at index %d does not match", actualStage.Category, idx)
					}
				}

				return nil
			}, "600s", "10s").Should(BeNil())
		})

		It("should deploy the service to the test environment", func() {
			type state struct {
				Name         string
				ActionName   string
				ActionStatus string
			}
			wantedStates := []state{
				{
					Name:         "Source",
					ActionName:   fmt.Sprintf("SourceCodeFor-%s", appName),
					ActionStatus: "Succeeded",
				},
				{
					Name:         "Build",
					ActionName:   "Build",
					ActionStatus: "Succeeded",
				},
				{
					Name:         "DeployTo-test",
					ActionName:   "CreateOrUpdate-frontend-test",
					ActionStatus: "Succeeded",
				},
			}

			Eventually(func() error {
				out, err := copilot.PipelineStatus(client.PipelineStatusInput{
					Name: testPipelineName,
				})
				if err != nil {
					return err
				}
				if len(wantedStates) != len(out.States) {
					return fmt.Errorf("len of pipeline states do not match: %v", out.States)
				}
				for idx, actualState := range out.States {
					if wantedStates[idx].Name != actualState.Name {
						return fmt.Errorf("state name %s at index %d does not match", actualState.Name, idx)
					}
					if len(actualState.Actions) != 1 {
						return fmt.Errorf("no action yet for state name %s", actualState.Name)
					}
					if wantedStates[idx].ActionName != actualState.Actions[0].Name {
						return fmt.Errorf("action name %v for state %s does not match at index %d", actualState.Actions[0], actualState.Name, idx)
					}
					if wantedStates[idx].ActionStatus != actualState.Actions[0].Status {
						return fmt.Errorf("action status %v for state %s does not match at index %d", actualState.Actions[0], actualState.Name, idx)
					}
				}
				return nil
			}, "1200s", "15s").Should(BeNil())
		})
	})

	Context("test pipeline service should be queryable post-release", func() {
		It("service should include a valid URL", func() {
			out, err := copilot.SvcShow(&client.SvcShowRequest{
				AppName: appName,
				Name:    "frontend",
			})
			Expect(err).NotTo(HaveOccurred())

			routes := make(map[string]string)
			for _, route := range out.Routes {
				routes[route.Environment] = route.URL
			}
			for _, env := range []string{"test"} {
				Eventually(func() (int, error) {
					resp, fetchErr := http.Get(routes[env])
					return resp.StatusCode, fetchErr
				}, "30s", "1s").Should(Equal(200))
			}
			// Check that the addons stack was created.
			Eventually(func() error {
				for _, variable := range out.Variables {
					if variable.Name == "E2EPIPELINEADDON_NAME" {
						return nil
					}
				}
				return fmt.Errorf("addons variable %s not found", "E2EPIPELINEADDON_NAME")
			})
		})
	})
	Context("when creating the prod pipeline stack", func() {
		It("checks out the prod git branch", func() {
			cmd := exec.Command("git", "checkout", "prod")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = repoName
			Expect(cmd.Run()).NotTo(HaveOccurred())
		})
		It("should start creating the pipeline stack", func() {
			_, err := copilot.PipelineDeploy(client.PipelineDeployInput{
				Name: prodPipelineName,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should show prod pipeline details once the stack is created", func() {
			type stage struct {
				Name     string
				Category string
			}
			wantedStages := []stage{
				{
					Name:     "Source",
					Category: "Source",
				},
				{
					Name:     "Build",
					Category: "Build",
				},
				{
					Name:     "DeployTo-prod",
					Category: "Deploy",
				},
			}

			Eventually(func() error {
				out, err := copilot.PipelineShow(client.PipelineShowInput{
					Name: prodPipelineName,
				})
				switch {
				case err != nil:
					return err
				case out.Name == "":
					return fmt.Errorf("pipeline name is empty: %v", out)
				case out.Name != prodPipelineName:
					return fmt.Errorf("expected pipeline name %q, got %q", prodPipelineName, out.Name)
				case len(out.Stages) != len(wantedStages):
					return fmt.Errorf("pipeline stages do not match: %v", out.Stages)
				}

				for idx, actualStage := range out.Stages {
					if wantedStages[idx].Name != actualStage.Name {
						return fmt.Errorf("stage name %s at index %d does not match", actualStage.Name, idx)
					}
					if wantedStages[idx].Category != actualStage.Category {
						return fmt.Errorf("stage category %s at index %d does not match", actualStage.Category, idx)
					}
				}
				return nil
			}, "600s", "10s").Should(BeNil())
		})

		It("should deploy the service to the prod environment", func() {
			type state struct {
				Name         string
				ActionName   string
				ActionStatus string
			}
			wantedStates := []state{
				{
					Name:         "Source",
					ActionName:   fmt.Sprintf("SourceCodeFor-%s", appName),
					ActionStatus: "Succeeded",
				},
				{
					Name:         "Build",
					ActionName:   "Build",
					ActionStatus: "Succeeded",
				},
				{
					Name:         "DeployTo-prod",
					ActionName:   "CreateOrUpdate-frontend-prod",
					ActionStatus: "Succeeded",
				},
			}

			Eventually(func() error {
				out, err := copilot.PipelineStatus(client.PipelineStatusInput{
					Name: prodPipelineName,
				})
				if err != nil {
					return err
				}
				if len(wantedStates) != len(out.States) {
					return fmt.Errorf("len of pipeline states do not match: %v", out.States)
				}
				for idx, actualState := range out.States {
					if wantedStates[idx].Name != actualState.Name {
						return fmt.Errorf("state name %s at index %d does not match", actualState.Name, idx)
					}
					if len(actualState.Actions) != 1 {
						return fmt.Errorf("no action yet for state name %s", actualState.Name)
					}
					if wantedStates[idx].ActionName != actualState.Actions[0].Name {
						return fmt.Errorf("action name %v for state %s does not match at index %d", actualState.Actions[0], actualState.Name, idx)
					}
					if wantedStates[idx].ActionStatus != actualState.Actions[0].Status {
						return fmt.Errorf("action status %v for state %s does not match at index %d", actualState.Actions[0], actualState.Name, idx)
					}
				}
				return nil
			}, "1200s", "15s").Should(BeNil())
		})
	})

	Context("prod pipeline service should be queryable post-release", func() {
		It("service should include a valid URL", func() {
			out, err := copilot.SvcShow(&client.SvcShowRequest{
				AppName: appName,
				Name:    "frontend",
			})
			Expect(err).NotTo(HaveOccurred())

			routes := make(map[string]string)
			for _, route := range out.Routes {
				routes[route.Environment] = route.URL
			}
			for _, env := range []string{"prod"} {
				Eventually(func() (int, error) {
					resp, fetchErr := http.Get(routes[env])
					return resp.StatusCode, fetchErr
				}, "30s", "1s").Should(Equal(200))
			}
			// Check that the addons stack was created.
			Eventually(func() error {
				for _, variable := range out.Variables {
					if variable.Name == "E2EPIPELINEADDON_NAME" {
						return nil
					}
				}
				return fmt.Errorf("addons variable %s not found", "E2EPIPELINEADDON_NAME")
			})
		})
	})
})
