package zapcore

import "go.uber.org/multierr"

type fastSynced struct {
	Core
}

func FastSync(c Core) Core {
	return &fastSynced{c}
}

func (f *fastSynced) Check(entry Entry, ce *CheckedEntry) *CheckedEntry {
	if f.Core.Enabled(entry.Level) {
		return ce.AddCore(entry, f)
	}
	return ce
}

func (f *fastSynced) Write(entry Entry, fields []Field) error {
	return multierr.Append(f.Core.Write(entry, fields), f.Core.Sync())
}
