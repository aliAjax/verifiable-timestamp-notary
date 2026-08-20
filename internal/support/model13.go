package support

// Model13 groups validation and lifecycle helpers used by the notary bounded contexts.
type Record13_1 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_1(id string) Record13_1 {
	return Record13_1{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_1) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_1) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_1) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_1) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_1) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_1) Label(k string) string { return r.Labels[k] }
func (r Record13_1) Clone() Record13_1 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_1) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_1) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_1) Key() string { return r.ID + ":" + r.Name }

type Record13_2 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_2(id string) Record13_2 {
	return Record13_2{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_2) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_2) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_2) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_2) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_2) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_2) Label(k string) string { return r.Labels[k] }
func (r Record13_2) Clone() Record13_2 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_2) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_2) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_2) Key() string { return r.ID + ":" + r.Name }

type Record13_3 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_3(id string) Record13_3 {
	return Record13_3{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_3) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_3) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_3) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_3) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_3) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_3) Label(k string) string { return r.Labels[k] }
func (r Record13_3) Clone() Record13_3 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_3) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_3) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_3) Key() string { return r.ID + ":" + r.Name }

type Record13_4 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_4(id string) Record13_4 {
	return Record13_4{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_4) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_4) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_4) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_4) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_4) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_4) Label(k string) string { return r.Labels[k] }
func (r Record13_4) Clone() Record13_4 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_4) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_4) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_4) Key() string { return r.ID + ":" + r.Name }

type Record13_5 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_5(id string) Record13_5 {
	return Record13_5{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_5) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_5) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_5) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_5) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_5) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_5) Label(k string) string { return r.Labels[k] }
func (r Record13_5) Clone() Record13_5 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_5) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_5) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_5) Key() string { return r.ID + ":" + r.Name }

type Record13_6 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_6(id string) Record13_6 {
	return Record13_6{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_6) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_6) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_6) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_6) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_6) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_6) Label(k string) string { return r.Labels[k] }
func (r Record13_6) Clone() Record13_6 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_6) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_6) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_6) Key() string { return r.ID + ":" + r.Name }

type Record13_7 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_7(id string) Record13_7 {
	return Record13_7{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_7) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_7) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_7) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_7) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_7) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_7) Label(k string) string { return r.Labels[k] }
func (r Record13_7) Clone() Record13_7 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_7) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_7) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_7) Key() string { return r.ID + ":" + r.Name }

type Record13_8 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_8(id string) Record13_8 {
	return Record13_8{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_8) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_8) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_8) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_8) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_8) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_8) Label(k string) string { return r.Labels[k] }
func (r Record13_8) Clone() Record13_8 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_8) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_8) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_8) Key() string { return r.ID + ":" + r.Name }

type Record13_9 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_9(id string) Record13_9 {
	return Record13_9{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_9) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_9) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_9) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_9) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_9) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_9) Label(k string) string { return r.Labels[k] }
func (r Record13_9) Clone() Record13_9 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_9) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_9) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_9) Key() string { return r.ID + ":" + r.Name }

type Record13_10 struct {
	ID      string
	Name    string
	State   string
	Version int64
	Labels  map[string]string
}

func NewRecord13_10(id string) Record13_10 {
	return Record13_10{ID: id, State: "pending", Version: 1, Labels: map[string]string{}}
}
func (r *Record13_10) Activate() {
	if r.State == "pending" {
		r.State = "active"
		r.Version++
	}
}
func (r *Record13_10) Suspend() {
	if r.State == "active" {
		r.State = "suspended"
		r.Version++
	}
}
func (r *Record13_10) Close() {
	if r.State != "closed" {
		r.State = "closed"
		r.Version++
	}
}
func (r Record13_10) Valid() bool { return r.ID != "" && r.Version > 0 && r.State != "" }
func (r *Record13_10) SetLabel(k, v string) {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	r.Labels[k] = v
}
func (r Record13_10) Label(k string) string { return r.Labels[k] }
func (r Record13_10) Clone() Record13_10 {
	c := r
	c.Labels = map[string]string{}
	for k, v := range r.Labels {
		c.Labels[k] = v
	}
	return c
}
func (r Record13_10) CanTransition(next string) bool {
	switch r.State {
	case "pending":
		return next == "active" || next == "closed"
	case "active":
		return next == "suspended" || next == "closed"
	case "suspended":
		return next == "active" || next == "closed"
	case "closed":
		return false
	}
	return false
}
func (r *Record13_10) Transition(next string) bool {
	if !r.CanTransition(next) {
		return false
	}
	r.State = next
	r.Version++
	return true
}
func (r Record13_10) Key() string { return r.ID + ":" + r.Name }
