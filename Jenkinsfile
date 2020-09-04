#!/usr/bin/env groovy
// Centralize Jenkins Pipelines configuration using Shared Libraries
// https://github.com/AfterShip/deployment-jenkins-pipeline-library-ci-aftership-org
@Library("jenkins-pipeline-library@automation") _
entry {
	isDryRun = false
	flow = "golang"
	configInfo = [
	      deploymentGroup          : "automizely",
	      appName                  : "deployment-lmstfy",
	      gitRepoName              : "deployment-lmstfy.git",
	      chartName                : "golang-http",
	      essentialDockerImage     : "golang-essential",
	      essentialTag             : "golang-1.13.5",
	      requireStaticAsset       : false,
	      uploadAssetCredential    : "",
          domainType               : "automizely",
	      unitTest                 : "",
	      integrationTest          : "",
	      useEnvironmentVariable   : true,
	      hasReleaseEnvironment    : false,
	      hasStagingEnvironment    : false,
	      hasProductionEnvironment : false,
	      useHelmReleaseShortName  : true
	    ]
}
