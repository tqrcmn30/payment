package controller

import "payment/services"

type ControllerManager struct {
	*EntityController
	*UserAccountController
	*PaymentTransactionController
}

func NewControllerManager(serviceManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		EntityController:             NewEntityController(*serviceManager),
		UserAccountController:        NewUserAccountController(*serviceManager),
		PaymentTransactionController: NewPaymentTransactionController(*serviceManager),
	}
}
