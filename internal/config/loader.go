package config

type Source interface{ Lookup(string) (string, bool) }
type MapSource map[string]string

func (m MapSource) Lookup(k string) (string, bool) { v, ok := m[k]; return v, ok }

type Loader struct{ Sources []Source }

func (l Loader) Get(key, def string) string {
	for _, s := range l.Sources {
		if v, ok := s.Lookup(key); ok && v != "" {
			return v
		}
	}
	return def
}
func (l Loader) Required(key string) (string, bool) { v := l.Get(key, ""); return v, v != "" }
