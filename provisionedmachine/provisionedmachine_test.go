package provisionedmachine_test

import (
	"testing"

	"github.com/platform9/ssh-provider/provisionedmachine"
	sshproviderv1 "github.com/platform9/ssh-provider/sshproviderconfig/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func TestNewFromConfigMap(t *testing.T) {
	cm := &corev1.ConfigMap{
		Data: map[string]string{
			"provisionedMachine": `
  sshConfig:
    host: "10.4.252.94"
    port: 22
    publicKeys:
    - "ssh-dss AAAAB3NzaC1kc3MAAACBAPkFvDBlzDXDt/R1vMyASCyKohERRWVq2KHtz500N6yy6H7VEgXwEA8bDmAN9Xye0GBqFc/WmfV757+y6vuPul/t8Re5APTLCFH/Q4XBp2kxchVGtyGB6ihVwMoGar2IMMQuneVkz0+/fn92fC6wOZtT+YIAMpWghieqJhUjtgz7AAAAFQDsH/tN4UZIpP0sfsAxwh/o76LB6QAAAIA8OAwT8ufxdXB7EW9T7oP3AEKuvJ95q+eSyBEPJsOfaWWzGLsvy1lpbpuTm/Ffz77mSKKQnFO4PiZgCX+OTdTXjgc1JN9v+GPEmlXA2FaA4cDsiKMyXdgN1ncNnlfa1ZVRwzg9mCqsPrX2Zt2r0o0N9LSB+aQ62WOPEdtRqNVp+AAAAIAZUUOUNg0Jd8vTtj0p2zY5AVcBBlB+v41keodjbgxkT1gwX8FvRGyHzyqIJr0Vbyj4i7Yj7hlLI6JgkxGLlDcZSO8lqLll133tE75XY70yMTP4ff8Dp18II3PDUg/fqrDSmC9ZfsXY3dYQG7pfBnz9EOziWp3EzkO1CSe0n6RKtg== root@localhost"
    - "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLd8oe2tsyRTuctnUbIRAQMzbe00+E9IlPDi7znwXf5uOmIPKiG641tWEfCw1HIJBz10jkxljqNG+4zbOjFnhqg= root@localhost"
    - "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMlvMz8+NI/SxOkQCmNZXFdDYuJ+sRgBLWCEWGPTloEn root@localhost"
    - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDc2MCBwSo9MONFrdmNV+ouzAZIY83f+S/SUQUuWuIHy5lmtXVkBJFSK/3meYZQVoxVIQKyUshRG+6lQ2lwOrPY3UoMiX7ezYX3d7Zb594cHJ7lxj38I0oRJpllYF6xjIebHG7Sl8OEBw/eWu94+ISGmTGokxviBBJ9Rhkic3/NE8Kf0K62WoXtiHiOdhCPw9vSz3j1w/7KdN0hKxzEmRsBxBiu4aOnBKTBAYyys0Mc6F6pfFkzEG/q14Qn2aje325mnS+FiKKZcxEZiHtrqe220jYbvsOxEsYw3Iz6jWsj5kz4EMYYFj0cSiLW2b7ioK6nsIYwBgLn1Wii9SiW5+ib root@localhost"
    secretName: ssh-creds
  vipNetworkInterface: "eth0"`,
		},
	}
	_, err := provisionedmachine.NewFromConfigMap(cm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestToConfigMap(t *testing.T) {
	pm := provisionedmachine.ProvisionedMachine{
		SSHConfig: &sshproviderv1.SSHConfig{
			Host: "10.4.252.94",
			Port: 22,
			PublicKeys: []string{
				"ssh-dss AAAAB3NzaC1kc3MAAACBAPkFvDBlzDXDt/R1vMyASCyKohERRWVq2KHtz500N6yy6H7VEgXwEA8bDmAN9Xye0GBqFc/WmfV757+y6vuPul/t8Re5APTLCFH/Q4XBp2kxchVGtyGB6ihVwMoGar2IMMQuneVkz0+/fn92fC6wOZtT+YIAMpWghieqJhUjtgz7AAAAFQDsH/tN4UZIpP0sfsAxwh/o76LB6QAAAIA8OAwT8ufxdXB7EW9T7oP3AEKuvJ95q+eSyBEPJsOfaWWzGLsvy1lpbpuTm/Ffz77mSKKQnFO4PiZgCX+OTdTXjgc1JN9v+GPEmlXA2FaA4cDsiKMyXdgN1ncNnlfa1ZVRwzg9mCqsPrX2Zt2r0o0N9LSB+aQ62WOPEdtRqNVp+AAAAIAZUUOUNg0Jd8vTtj0p2zY5AVcBBlB+v41keodjbgxkT1gwX8FvRGyHzyqIJr0Vbyj4i7Yj7hlLI6JgkxGLlDcZSO8lqLll133tE75XY70yMTP4ff8Dp18II3PDUg/fqrDSmC9ZfsXY3dYQG7pfBnz9EOziWp3EzkO1CSe0n6RKtg== root@localhost",
				"ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLd8oe2tsyRTuctnUbIRAQMzbe00+E9IlPDi7znwXf5uOmIPKiG641tWEfCw1HIJBz10jkxljqNG+4zbOjFnhqg= root@localhost",
				"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMlvMz8+NI/SxOkQCmNZXFdDYuJ+sRgBLWCEWGPTloEn root@localhost",
				"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDc2MCBwSo9MONFrdmNV+ouzAZIY83f+S/SUQUuWuIHy5lmtXVkBJFSK/3meYZQVoxVIQKyUshRG+6lQ2lwOrPY3UoMiX7ezYX3d7Zb594cHJ7lxj38I0oRJpllYF6xjIebHG7Sl8OEBw/eWu94+ISGmTGokxviBBJ9Rhkic3/NE8Kf0K62WoXtiHiOdhCPw9vSz3j1w/7KdN0hKxzEmRsBxBiu4aOnBKTBAYyys0Mc6F6pfFkzEG/q14Qn2aje325mnS+FiKKZcxEZiHtrqe220jYbvsOxEsYw3Iz6jWsj5kz4EMYYFj0cSiLW2b7ioK6nsIYwBgLn1Wii9SiW5+ib root@localhost",
			},
			SecretName: "ssh-creds",
		},
		VIPNetworkInterface: "eth0",
	}

	cm := &corev1.ConfigMap{
		Data: make(map[string]string),
	}
	err := pm.ToConfigMap(cm)
	if err != nil {
		t.Fatal(err)
	}
}
