package model

type Event_Instansi struct {
	Id                string `json:"id" form:"id"`
	Event_id          string `json:"event_id" form:"event_id"`
	Instansi_id       string `json:"instansi_id" form:"instansi_id"`
	Instansi_nama     string `json:"instansi_nama" form:"instansi_nama"`
	Satuan_kerja_id   string `json:"satuan_kerja_id" form:"satuan_kerja_id"`
	Satuan_kerja_nama string `json:"satuan_kerja_nama" form:"satuan_kerja_nama"`
	Jenis             string `json:"jenis" form:"jenis"`
	Jenis_instansi_id string `json:"jenis_instansi_id" form:"jenis_instansi_id"`
	Status_aktif      bool   `json:"status_aktif" form:"status_aktif"`
	Kanreg_id         string `json:"kanreg_id" form:"kanreg_id"`
	Kanreg_nama       string `json:"kanreg_nama" form:"kanreg_nama"`
	Periode           int    `json:"periode" form:"periode"`
	StatusPeriode     string `json:"status_periode" form:"status_periode" sql:"-"`
}

func (Event_Instansi) TableName() string {
	return "event_instansi"
}
