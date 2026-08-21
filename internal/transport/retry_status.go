package transport

func RetryableStatus(code int) bool { return code >= 500 }
func RetryBudget(code int) int { if RetryableStatus(code) { return 3 }; return 0 }
