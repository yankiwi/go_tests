package definitions

type Device map[string]interface{}

var ArubaInstant = Device{}

func init() {
  ArubaInstant["versionCommand"] = "show version"
  ArubaInstant["versionRegex"] = "^.*$"
  ArubaInstant["memoryCommand"] = "show memory"
  ArubaInstant["memoryCommonRegex"] = "^.*total: (.*) used: (.*) total: (.*)$"
  ArubaInstant["memoryTotalRegex"] = ""
  ArubaInstant["memoryUsedRegex"] = ""
  ArubaInstant["memoryFreeRegex"] = ""
  ArubaInstant["memoryTotal"] = Resource{regex: "memoryCommonRegex", position: 1}
  ArubaInstant["memoryUsed"] = Resource{regex: "memoryCommonRegex", position: 2}
  ArubaInstant["memoryFree"] = Resource{regex: "memoryCommonRegex", position: 3}
  ArubaInstant["cpuCommand"] = "show cpu"
  ArubaInstant["cpuCommonRegEx"] = ""
  ArubaInstant["cpuUsedRegex"] = "^.*used (.*)$"
  ArubaInstant["cpuIdleRegex"] = "^.*idle (.*)$"
  ArubaInstant["cpuUsed"] = Resource{regex: "cpuUsedRegex", position: 1}
  ArubaInstant["cpuIdle"] = Resource{regex: "cpuIdleRegex", position: 1}

}

