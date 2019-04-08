package models

type BaseAttributes struct {
	DistinguishedName string
	Status            string
	Description       string
	ClassName         string
}

func (ba *BaseAttributes) ToMap() (map[string]string, error) {
	result := make(map[string]string)
	result["dn"] = ba.DistinguishedName
	result["descr"] = ba.Description
	result["status"] = ba.Status
	result["classname"] = ba.ClassName
	return result, nil
}
