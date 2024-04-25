// Copyright The TBox Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

var (
	// DomainThis specifies generic domain [general purpose domain]
	DomainThis = NewDomain("this")
	// DomainSrc specifies abstract source [general purpose domain]
	DomainSrc = NewDomain("src")
	// DomainDst specifies abstract destination [general purpose domain]
	DomainDst = NewDomain("dst")
	// DomainReference specifies abstract reference [general purpose domain]
	DomainReference = NewDomain("reference")
	// DomainContext specifies abstract context [general purpose domain]
	DomainContext = NewDomain("context")
	// DomainTask specifies abstract task [general purpose domain]
	DomainTask = NewDomain("task")
	// DomainTaskResult specifies abstract task result [general purpose domain]
	DomainTaskResult = NewDomain("task_result")
	// DomainFile specifies abstract file [general purpose domain]
	DomainFile = NewDomain("file")
	// DomainStatus specifies abstract status [general purpose domain]
	DomainStatus = NewDomain("status")
	// DomainParent specifies abstract parent [general purpose domain]
	DomainParent = NewDomain("parent")
	// DomainReport specifies abstract report [general purpose domain]
	DomainReport = NewDomain("report")
	// DomainResult specifies abstract result [general purpose domain]
	DomainResult = NewDomain("result")
	// DomainInterim specifies abstract interim entities [general purpose domain]
	DomainInterim = NewDomain("interim")
	// DomainAddress specifies abstract address [general purpose domain]
	DomainAddress = NewDomain("address")
	// DomainUser specifies abstract user address [general purpose domain]
	DomainUser = NewDomain("user")
	// DomainProject specifies abstract project address [general purpose domain]
	DomainProject = NewDomain("project")
	// DomainAsset specifies abstract asset address [general purpose domain]
	DomainAsset = NewDomain("asset")

	// DomainS3 specifies S3 domain [predefined address domain]
	DomainS3 = NewDomain("s3")
	// DomainKafka specifies Kafka domain [predefined address domain]
	DomainKafka = NewDomain("kafka")
	// DomainDigest specifies digest (hash, etc.) [predefined address domain]
	DomainDigest = NewDomain("digest")
	// DomainUUID specifies UUID [predefined address domain]
	DomainUUID = NewDomain("uuid")
	// DomainUserID specifies user ID [predefined address domain]
	DomainUserID = NewDomain("user_id")
	// DomainDirname specifies name of the directory [predefined address domain]
	DomainDirname = NewDomain("dirname")
	// DomainFilename specifies file name [predefined address domain]
	DomainFilename = NewDomain("filename")
	// DomainURL specifies URL [predefined address domain]
	DomainURL = NewDomain("url")
	// DomainDomain specifies domain [predefined address domain]
	DomainDomain = NewDomain("domain")
	// DomainMachineID specifies machine ID [predefined address domain]
	DomainMachineID = NewDomain("machine_id")
	// DomainAssetID specifies asset ID [predefined address domain]
	DomainAssetID = NewDomain("asset_id")
	// DomainTaskID specifies task ID [predefined address domain]
	DomainTaskID = NewDomain("task_id")
	// DomainProjectID specifies project ID [predefined address domain]
	DomainProjectID = NewDomain("project_id")
	// DomainEmail specifies email [predefined address domain]
	DomainEmail = NewDomain("email")
	// DomainCustom specifies custom domain [predefined address domain]
	DomainCustom = NewDomain("custom")

	// Domains specifies list of all registered domains
	Domains = []*Domain{
		// General purpose domains
		DomainThis,
		DomainSrc,
		DomainDst,
		DomainReference,
		DomainContext,
		DomainTask,
		DomainTaskResult,
		DomainFile,
		DomainStatus,
		DomainParent,
		DomainReport,
		DomainResult,
		DomainInterim,
		DomainAddress,
		DomainUser,
		DomainProject,
		DomainAsset,
		// Predefined address domains
		DomainS3,
		DomainKafka,
		DomainDigest,
		DomainUUID,
		DomainUserID,
		DomainDirname,
		DomainFilename,
		DomainURL,
		DomainDomain,
		DomainMachineID,
		DomainAssetID,
		DomainTaskID,
		DomainProjectID,
		DomainEmail,
		DomainCustom,
	}
)

// RegisterDomain tries to register specified Domain.
// Domain must be non-equal to all registered domains.
// Returns nil in case Domain can not be registered, say it is equal to previously registered domain
func RegisterDomain(domain *Domain) *Domain {
	if FindDomain(domain) != nil {
		// Such a domain already exists
		return nil
	}
	Domains = append(Domains, domain)
	return domain
}

// MustRegisterDomain the same as RegisterDomain but with panic
func MustRegisterDomain(domain *Domain) {
	if RegisterDomain(domain) == nil {
		panic("unable to register domain")
	}
}

// FindDomain returns registered domain with the same string value as provided
func FindDomain(domain *Domain) *Domain {
	return DomainFromString(domain.Name)
}

// NormalizeDomain returns either registered domain with the same string value as provided domain or provided domain.
func NormalizeDomain(domain *Domain) *Domain {
	if f := FindDomain(domain); f != nil {
		// Return registered domain
		return f
	}
	// Unable to find registered domain, return provided domain
	return domain
}

// DomainFromString tries to find registered domain with specified string value
func DomainFromString(str string) *Domain {
	d := NewDomain().SetName(str)
	for _, domain := range Domains {
		if domain.Equals(d) {
			return domain
		}
	}
	return nil
}
