// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package common

import (
	"fmt"
	"net"
)

func NewRetinaEndpoint(name, namespace string, ips *IPAddresses) *RetinaEndpoint {
	return &RetinaEndpoint{
		BaseObject: GetBaseObject(name, namespace, ips),
	}
}

func (ep *RetinaEndpoint) DeepCopy() interface{} {
	ep.RLock()
	defer ep.RUnlock()

	newEp := &RetinaEndpoint{
		BaseObject: ep.BaseObject.DeepCopy(),
	}

	if ep.ownerRefs != nil {
		newEp.ownerRefs = make([]*OwnerReference, len(ep.ownerRefs))
		for i, ownerRef := range ep.ownerRefs {
			newEp.ownerRefs[i] = ownerRef.DeepCopy()
		}
	}

	if ep.containers != nil {
		newEp.containers = make([]*RetinaContainer, len(ep.containers))
		for i, container := range ep.containers {
			newEp.containers[i] = container.DeepCopy()
		}
	}

	if ep.labels != nil {
		newEp.labels = make(map[string]string)
		for k, v := range ep.labels {
			newEp.labels[k] = v
		}
	}

	if ep.annotations != nil {
		newEp.annotations = make(map[string]string)
		for k, v := range ep.annotations {
			newEp.annotations[k] = v
		}
	}

	return newEp
}

func (ep *RetinaEndpoint) IPs() *IPAddresses {
	ep.RLock()
	defer ep.RUnlock()

	return ep.ips
}

func (ep *RetinaEndpoint) SetIPs(ips *IPAddresses) {
	ep.Lock()
	defer ep.Unlock()
	ep.ips = ips
}

func (ep *RetinaEndpoint) OwnerRefs() []*OwnerReference {
	ep.RLock()
	defer ep.RUnlock()
	return ep.ownerRefs
}

func (ep *RetinaEndpoint) SetOwnerRefs(ownerRefs []*OwnerReference) {
	ep.Lock()
	defer ep.Unlock()
	ep.ownerRefs = ownerRefs
}

func (ep *RetinaEndpoint) Containers() []*RetinaContainer {
	ep.RLock()
	defer ep.RUnlock()
	return ep.containers
}

func (ep *RetinaEndpoint) SetContainers(containers []*RetinaContainer) {
	ep.Lock()
	defer ep.Unlock()
	ep.containers = containers
}

func (ep *RetinaEndpoint) Labels() map[string]string {
	ep.RLock()
	defer ep.RUnlock()
	return ep.labels
}

func (ep *RetinaEndpoint) FormattedLabels() []string {
	ep.RLock()
	defer ep.RUnlock()
	labels := make([]string, 0)
	for k, v := range ep.labels {
		labels = append(labels, fmt.Sprintf("%s=%s", k, v))
	}

	return labels
}

func (ep *RetinaEndpoint) SetLabels(labels map[string]string) {
	ep.Lock()
	defer ep.Unlock()
	ep.labels = labels
}

func (ep *RetinaEndpoint) Annotations() map[string]string {
	ep.RLock()
	defer ep.RUnlock()
	return ep.annotations
}

func (ep *RetinaEndpoint) SetAnnotations(annotations map[string]string) {
	ep.Lock()
	defer ep.Unlock()
	if annotations != nil {
		ep.annotations = make(map[string]string)
		if _, ok := annotations[RetinaPodAnnotation]; ok {
			ep.annotations[RetinaPodAnnotation] = annotations[RetinaPodAnnotation]
		}
	}
}

func (e *RetinaEndpoint) PrimaryIP() (string, error) {
	e.RLock()
	defer e.RUnlock()

	if e.ips != nil {
		pip := e.ips.PrimaryIP()
		if pip != "" {
			return pip, nil
		}
	}

	return "", fmt.Errorf("no primary IP found for endpoint %s", e.Key())
}

func (e *RetinaEndpoint) PrimaryNetIP() (net.IP, error) {
	e.RLock()
	defer e.RUnlock()

	if e.ips != nil {
		pip := e.ips.PrimaryNetIP()
		if pip != nil {
			return pip, nil
		}
	}

	return nil, fmt.Errorf("no primary IP found for endpoint %s", e.Key())
}

func (o *OwnerReference) DeepCopy() *OwnerReference {
	return &OwnerReference{
		Kind:       o.Kind,
		Name:       o.Name,
		UID:        o.UID,
		Controller: o.Controller,
	}
}

func (c *RetinaContainer) DeepCopy() *RetinaContainer {
	return &RetinaContainer{
		Name: c.Name,
		ID:   c.ID,
	}
}
