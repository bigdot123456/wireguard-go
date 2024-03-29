/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2019 WireGuard LLC. All Rights Reserved.
 */

package device

import (
	"encoding/binary"
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	sockoptIP_UNICAST_IF   = 31
	sockoptIPV6_UNICAST_IF = 31
)

func (device *Device) BindSocketToInterface4(interfaceIndex uint32) error {
	/* MSDN says for IPv4 this needs to be in net byte order, so that it's like an IP address with leading zeros. */
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, interfaceIndex)
	interfaceIndex = *(*uint32)(unsafe.Pointer(&bytes[0]))

	if device.net.bind == nil {
		return errors.New("Bind is not yet initialized")
	}

	sysconn, err := device.net.bind.(*nativeBind).ipv4.SyscallConn()
	if err != nil {
		return err
	}
	err2 := sysconn.Control(func(fd uintptr) {
		err = windows.SetsockoptInt(windows.Handle(fd), windows.IPPROTO_IP, sockoptIP_UNICAST_IF, int(interfaceIndex))
	})
	if err2 != nil {
		return err2
	}
	if err != nil {
		return err
	}
	return nil
}

func (device *Device) BindSocketToInterface6(interfaceIndex uint32) error {
	sysconn, err := device.net.bind.(*nativeBind).ipv6.SyscallConn()
	if err != nil {
		return err
	}
	err2 := sysconn.Control(func(fd uintptr) {
		err = windows.SetsockoptInt(windows.Handle(fd), windows.IPPROTO_IPV6, sockoptIPV6_UNICAST_IF, int(interfaceIndex))
	})
	if err2 != nil {
		return err2
	}
	if err != nil {
		return err
	}
	return nil
}
