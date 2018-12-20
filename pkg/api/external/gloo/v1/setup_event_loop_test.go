// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("SetupEventLoop", func() {
	var (
		namespace string
		emitter   SetupEmitter
		err       error
	)

	BeforeEach(func() {

		settingsClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		settingsClient, err := NewSettingsClient(settingsClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewSetupEmitter(settingsClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Settings().Write(NewSettings(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockSetupSyncer{}
		el := NewSetupEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(sync.Synced, 5*time.Second).Should(BeTrue())
	})
})

type mockSetupSyncer struct {
	synced bool
	mutex  sync.Mutex
}

func (s *mockSetupSyncer) Synced() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.synced
}

func (s *mockSetupSyncer) Sync(ctx context.Context, snap *SetupSnapshot) error {
	s.mutex.Lock()
	s.synced = true
	s.mutex.Unlock()
	return nil
}
