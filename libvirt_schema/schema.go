// Copyright 2017 Kumina, https://kumina.nl/
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libvirt_schema

//Domain libvirt xml structure
type Domain struct {
	Devices  Devices  `xml:"devices"`
	Metadata Metadata `xml:"metadata"`
}

//Metadata xml structure of domain
type Metadata struct {
	NovaInstance NovaInstance `xml:"instance"`
}

//NovaInstance xml structure of domain
type NovaInstance struct {
	InstanceName string `xml:"name"`
}

//Devices xml structure of domain
type Devices struct {
	Disks      []Disk      `xml:"disk"`
	Interfaces []Interface `xml:"interface"`
}

//Disk xml structure of domain
type Disk struct {
	Source DiskSource `xml:"source"`
	Target DiskTarget `xml:"target"`
}

//DiskSource xml structure of domain
type DiskSource struct {
	File string `xml:"file,attr"`
}

//DiskTarget xml structure of domain
type DiskTarget struct {
	Device string `xml:"dev,attr"`
}

//Interface xml structure of domain
type Interface struct {
	Source InterfaceSource `xml:"source"`
	Target InterfaceTarget `xml:"target"`
}

//InterfaceSource xml structure of domain
type InterfaceSource struct {
	Bridge string `xml:"bridge,attr"`
}

//InterfaceTarget xml structure of domain
type InterfaceTarget struct {
	Device string `xml:"dev,attr"`
}
