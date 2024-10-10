package hooks

import "os"

/*
 * This file is only ever generated once on the first generation and then is free to be modified.
 * Any hooks you wish to add should be registered in the InitHooks function. Feel free to define them
 * in this file or in separate files in the hooks package.
 */

func initHooks(h *Hooks) {
	h.registerBeforeRequestHook(&UserAgentPreRequestHook{})

	h.registerBeforeRequestHook(&APIURLRequestHook{})

	h.registerBeforeRequestHook(&HTTPDumpRequestHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_REQUEST") == "true",
	})

	h.registerAfterSuccessHook(&HTTPDumpResponseHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_RESPONSE") == "true",
	})
	h.registerAfterErrorHook(&HTTPDumpResponseHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_RESPONSE") == "true",
	})
}
