package status

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Conditions []v1.Condition

func (c *Conditions) Set(conditionType string, conditionStatus v1.ConditionStatus, conditionMessage string) {
	for i := range *c {
		condition := &[]v1.Condition(*c)[i]
		if condition.Type == conditionType {
			if condition.Status != conditionStatus || conditionMessage != conditionMessage {
				condition.LastTransitionTime = v1.Now()
				condition.ObservedGeneration = condition.ObservedGeneration + 1
			}
			condition.Message = conditionMessage
			condition.Status = conditionStatus
			return
		}
	}
	*c = append(*c, v1.Condition{
		Type:               conditionType,
		Status:             conditionStatus,
		Message:            conditionMessage,
		LastTransitionTime: v1.Now(),
	})
}

func (c *Conditions) Get(conditionType string) *v1.Condition {
	for _, condition := range *c {
		if condition.Type == conditionType {
			return &condition
		}
	}
	return nil
}
