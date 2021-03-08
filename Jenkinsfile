#!/usr/bin/env groovy
// Centralize Jenkins Pipelines configuration using Shared Libraries
// https://github.com/AfterShip/deployment-jenkins-pipeline-library-ci-aftership-org
@Library("jenkins-pipeline-library@automation") _
entry {
	isDryRun = false
	flow = "golang"
	configInfo = [
	      deploymentGroup          : "aftership",
          appName                  : "lmstfy.automizelyapi.org_main",
	      gitRepoName              : "lmstfy.automizelyapi.org_main.git",
	      chartName                : "golang-http",
	      essentialDockerImage     : "golang-essential",
	      essentialTag             : "golang-1.15.3",
	      requireStaticAsset       : false,
	      uploadAssetCredential    : "",
          domainType               : "automizelyapi.org",
	      unitTest                 : "",
	      integrationTest          : "",
	      useEnvironmentVariable   : true,
	      hasReleaseEnvironment    : false,
	      hasStagingEnvironment    : false,
	      hasProductionEnvironment : false,
	      useHelmReleaseShortName  : true

	    ]
}