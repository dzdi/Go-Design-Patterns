package state17

import "fmt"

// 状态机
type Machine struct {
	state IState
}

func (m2 *Machine) GetStateName() string {
	return m2.state.GetName()
}

func (m2 *Machine) SetState(state IState) {
	m2.state = state
}

func (m2 *Machine) Approval() {
	m2.state.Approval(m2)
}

func (m2 *Machine) Reject() {
	m2.state.Reject(m2)
}

type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

type leaderApproveState struct {
}

func (l leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

func (l leaderApproveState) Reject(m *Machine) {}

func (l leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

type financeApproveState struct {
}

func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
