package status

import (
	"github.com/samber/lo"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Conditions []v1.Condition

func (c *Conditions) Set(conditionType string, conditionStatus v1.ConditionStatus, conditionMessage string) {
	condition := c.Get(conditionType)
	if condition == nil {
		*c = append(*c, v1.Condition{})
	}
	*condition = v1.Condition{
		Type:               conditionType,
		Status:             conditionStatus,
		Message:            conditionMessage,
		ObservedGeneration: condition.ObservedGeneration + 1,
		LastTransitionTime: lo.Ternary(condition.Status == conditionStatus, condition.LastTransitionTime, v1.Now()),
	}
}

func (c *Conditions) Get(conditionType string) *v1.Condition {
	for _, condition := range *c {
		if condition.Type == conditionType {
			return &condition
		}
	}
	return nil
}
