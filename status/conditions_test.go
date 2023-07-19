package status_test

import (
	"testing"
	"time"

	"github.com/ellistarn/oper8/status"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite")
}

type TestStatus struct {
	Conditions []v1.Condition
}

var _ = Describe("Conditions", func() {
	It("should correctly toggle conditions", func() {
		testStatus := TestStatus{}
		// Condition is not set
		conditions := status.Conditions(testStatus.Conditions)
		Expect(conditions.Get("foo")).To(BeNil())
		Expect(testStatus).To(Equal(TestStatus{}))
		// Update the condition
		conditions.Set("foo", v1.ConditionTrue, "success")
		fooCondition := conditions.Get("foo")
		Expect(fooCondition.Type).To(Equal("foo"))
		Expect(fooCondition.Status).To(Equal(v1.ConditionTrue))
		Expect(fooCondition.Message).To(Equal("success"))
		Expect(fooCondition.LastTransitionTime.UnixNano()).To(BeNumerically(">", 0))
		Expect(fooCondition.ObservedGeneration).To(BeEquivalentTo(0))
		time.Sleep(1 * time.Nanosecond)
		// Update another condition
		conditions.Set("bar", v1.ConditionTrue, "success")
		barCondition := conditions.Get("bar")
		Expect(barCondition.Type).To(Equal("bar"))
		Expect(barCondition.Status).To(Equal(v1.ConditionTrue))
		Expect(barCondition.Message).To(Equal("success"))
		Expect(barCondition.LastTransitionTime.UnixNano()).To(BeNumerically(">", 0))
		Expect(barCondition.ObservedGeneration).To(BeEquivalentTo(0))
		time.Sleep(1 * time.Nanosecond)
		// transition the condition
		conditions.Set("foo", v1.ConditionFalse, "failed")
		updatedFooCondition := conditions.Get("foo")
		Expect(updatedFooCondition.Type).To(Equal("foo"))
		Expect(updatedFooCondition.Status).To(Equal(v1.ConditionFalse))
		Expect(updatedFooCondition.Message).To(Equal("failed"))
		Expect(updatedFooCondition.LastTransitionTime.UnixNano()).To(BeNumerically(">", fooCondition.LastTransitionTime.UnixNano()))
		Expect(updatedFooCondition.ObservedGeneration).To(BeEquivalentTo(1))
		time.Sleep(1 * time.Nanosecond)
		// Don't transition if the status is the same
		conditions.Set("bar", v1.ConditionTrue, "success")
		updatedBarCondition := conditions.Get("bar")
		Expect(updatedBarCondition.Type).To(Equal("bar"))
		Expect(updatedBarCondition.Status).To(Equal(v1.ConditionTrue))
		Expect(updatedBarCondition.Message).To(Equal("success"))
		Expect(updatedBarCondition.LastTransitionTime.UnixNano()).To(BeNumerically("==", barCondition.LastTransitionTime.UnixNano()))
		Expect(updatedBarCondition.ObservedGeneration).To(BeEquivalentTo(0))
	})
})
