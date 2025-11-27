# GpuAvailabilityPrototype

**Properties**

| Name          | Type                                    | Required | Description                                               |
| :------------ | :-------------------------------------- | :------- | :-------------------------------------------------------- |
| GpuClasses    | []string                                | ✅       | A list of available GPU class names                       |
| Cpu           | int64                                   | ❌       | The number of available CPU cores                         |
| Memory        | int64                                   | ❌       | The amount of available memory in MB                      |
| StorageAmount | int64                                   | ❌       | The amount of available storage in bytes                  |
| CountryCodes  | [][shared.CountryCode](country_code.md) | ❌       | A list of country codes where the resources are available |
