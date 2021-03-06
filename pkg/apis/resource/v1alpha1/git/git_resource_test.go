/*
Copyright 2019-2020 The Tekton Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package git_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	pipelinev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	resourcev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1/git"
	tb "github.com/tektoncd/pipeline/test/builder"
	"github.com/tektoncd/pipeline/test/names"
	corev1 "k8s.io/api/core/v1"
)

func TestNewGitResource_Invalid(t *testing.T) {
	if _, err := git.NewResource("override-with-git:latest", tb.PipelineResource("git-resource", tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGCS))); err == nil {
		t.Error("Expected error creating Git resource")
	}
}

func TestNewGitResource_Valid(t *testing.T) {
	for _, tc := range []struct {
		desc             string
		pipelineResource *resourcev1alpha1.PipelineResource
		want             *git.Resource
	}{{
		desc: "With Revision",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "Without Revision",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "With Refspec",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Refspec", "refs/changes/22/222134"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "refs/changes/22/222134",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "Without Refspec",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "With Submodules",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "Without Submodules",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Submodules", "false"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "With positive depth",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Depth", "8"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      8,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "With zero depth",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Depth", "0"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      0,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "Without SSLVerify",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Depth", "0"),
				tb.PipelineResourceSpecParam("SSLVerify", "false"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      0,
			SSLVerify:  false,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "With HTTPProxy",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Depth", "0"),
				tb.PipelineResourceSpecParam("HTTPProxy", "http-proxy.git.com"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      0,
			SSLVerify:  true,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "",
			NOProxy:    "",
		},
	}, {
		desc: "With HTTPSProxy",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Depth", "0"),
				tb.PipelineResourceSpecParam("HTTPSProxy", "https-proxy.git.com"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      0,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "",
		},
	}, {
		desc: "With NOProxy",
		pipelineResource: tb.PipelineResource("git-resource",
			tb.PipelineResourceSpec(resourcev1alpha1.PipelineResourceTypeGit,
				tb.PipelineResourceSpecParam("URL", "git@github.com:test/test.git"),
				tb.PipelineResourceSpecParam("Revision", "test"),
				tb.PipelineResourceSpecParam("Depth", "0"),
				tb.PipelineResourceSpecParam("NOProxy", "*"),
			),
		),
		want: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "test",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      0,
			SSLVerify:  true,
			HTTPProxy:  "",
			HTTPSProxy: "",
			NOProxy:    "*",
		},
	}} {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := git.NewResource("override-with-git:latest", tc.pipelineResource)
			if err != nil {
				t.Fatalf("Unexpected error creating Git resource: %s", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Mismatch of Git resource: %s", diff)
			}
		})
	}
}

func TestGitResource_Replacements(t *testing.T) {
	r := &git.Resource{
		Name:       "git-resource",
		Type:       resourcev1alpha1.PipelineResourceTypeGit,
		URL:        "git@github.com:test/test.git",
		Revision:   "master",
		Refspec:    "",
		Submodules: false,
		Depth:      16,
		SSLVerify:  false,
		HTTPProxy:  "http-proxy.git.com",
		HTTPSProxy: "https-proxy.git.com",
		NOProxy:    "*",
	}

	want := map[string]string{
		"name":       "git-resource",
		"type":       string(resourcev1alpha1.PipelineResourceTypeGit),
		"url":        "git@github.com:test/test.git",
		"revision":   "master",
		"refspec":    "",
		"submodules": "false",
		"depth":      "16",
		"sslVerify":  "false",
		"httpProxy":  "http-proxy.git.com",
		"httpsProxy": "https-proxy.git.com",
		"noProxy":    "*",
	}

	got := r.Replacements()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch of GitResource Replacements: %s", diff)
	}
}

func TestGitResource_GetDownloadTaskModifier(t *testing.T) {
	names.TestingSeed()

	for _, tc := range []struct {
		desc        string
		gitResource *git.Resource
		want        corev1.Container
	}{{
		desc: "With basic values",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-9l9zj",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}, {
		desc: "Without submodules",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-mz4c7",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-submodules=false",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}, {
		desc: "With more depth",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: true,
			Depth:      8,
			SSLVerify:  true,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-mssqb",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-depth",
				"8",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}, {
		desc: "Without sslVerify",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  false,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-78c5n",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-submodules=false",
				"-sslVerify=false",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}, {
		desc: "Without httpProxy",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  false,
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-6nl7g",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-submodules=false",
				"-sslVerify=false",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}, {
		desc: "Without httpsProxy",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  false,
			HTTPProxy:  "http-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-j2tds",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-submodules=false",
				"-sslVerify=false",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}, {
		desc: "Without noProxy",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  false,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "https-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-vr6ds",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-submodules=false",
				"-sslVerify=false",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
			},
		},
	}, {
		desc: "With Refspec",
		gitResource: &git.Resource{
			Name:       "git-resource",
			Type:       resourcev1alpha1.PipelineResourceTypeGit,
			URL:        "git@github.com:test/test.git",
			Revision:   "master",
			Refspec:    "refs/tags/v1.0:refs/tags/v1.0 refs/heads/master:refs/heads/master",
			GitImage:   "override-with-git:latest",
			Submodules: false,
			Depth:      1,
			SSLVerify:  true,
			HTTPProxy:  "http-proxy.git.com",
			HTTPSProxy: "https-proxy.git.com",
			NOProxy:    "no-proxy.git.com",
		},
		want: corev1.Container{
			Name:    "git-source-git-resource-l22wn",
			Image:   "override-with-git:latest",
			Command: []string{"/ko-app/git-init"},
			Args: []string{
				"-url",
				"git@github.com:test/test.git",
				"-revision",
				"master",
				"-path",
				"/test/test",
				"-refspec",
				"refs/tags/v1.0:refs/tags/v1.0 refs/heads/master:refs/heads/master",
				"-submodules=false",
			},
			WorkingDir: "/workspace",
			Env: []corev1.EnvVar{
				{Name: "TEKTON_RESOURCE_NAME", Value: "git-resource"},
				{Name: "HTTP_PROXY", Value: "http-proxy.git.com"},
				{Name: "HTTPS_PROXY", Value: "https-proxy.git.com"},
				{Name: "NO_PROXY", Value: "no-proxy.git.com"},
			},
		},
	}} {
		t.Run(tc.desc, func(t *testing.T) {
			ts := pipelinev1alpha1.TaskSpec{}
			modifier, err := tc.gitResource.GetInputTaskModifier(&ts, "/test/test")
			if err != nil {
				t.Fatalf("Unexpected error getting GetDownloadTaskModifier: %s", err)
			}

			if diff := cmp.Diff([]pipelinev1alpha1.Step{{Container: tc.want}}, modifier.GetStepsToPrepend()); diff != "" {
				t.Errorf("Mismatch of GitResource DownloadContainerSpec: %s", diff)
			}
		})
	}
}
