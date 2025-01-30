package bearden

func Create(constructor func() *ModuleFactory) *BearDenApplication {
	return &BearDenApplication{}
}
