package peremajaan

import (
	"code-echo/db"
	"code-echo/model"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)


func SimpanEvent(c echo.Context) error {
	//Inisiasi Model Baru dari Struct Even Instansi
	ModelEventInstansi := new(model.Event_Instansi)

	//Memperoleh isi input / parameter yang di kirim dari client
	if err := c.Bind(ModelEventInstansi); err != nil {
		return err
	}

	//Inisiasi Koneksi Database
	DbConn := db.Manager()

	//Generate UUID Untuk Id record yang baru
	u1 := uuid.Must(uuid.NewV4())
	Id := u1.String()

	//Memberi Nilai Id yang ada dalam ModelEventInstansi
	ModelEventInstansi.Id = Id

	// Insert Di tabel usulan
	 dbc := DbConn.Debug().Create(&ModelEventInstansi);

	 //Kondisi Jika Terjadi error saat eksekusi query
	 if dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Mereturn Response dalam format JSON
	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message":  "Berhasil Disimpan",
		"Id": Id,
	})
}



func UpdateEvent(c echo.Context) error {
	//Inisiasi Model Baru dari Struct Even Instansi
	ModelEventInstansi := new(model.Event_Instansi)

	//Memperoleh isi input / parameter yang di kirim dari client
	if err := c.Bind(ModelEventInstansi); err != nil {
		return err
	}

	//Inisiasi Koneksi Database
	DbConn := db.Manager()
	// Insert Di tabel usulan
	dbc := DbConn.Debug().Model(&ModelEventInstansi).Where("id = ?", ModelEventInstansi.Id).Update(ModelEventInstansi);

	//Kondisi Jika Terjadi error saat eksekusi query
	if dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Mereturn Response dalam format JSON
	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Berhasil Diupdate",
	})

}

func GetEvent(c echo.Context) error {
	DbCon := db.Manager()
	Event := []model.Event_Instansi{}
	DbCon.Find(&Event)
	for i := range Event{
		if Event[i].Periode == 0 {
			Event[i].StatusPeriode = "Periode Nonaktif"
		}
	}
	return c.JSON(http.StatusOK, Event)
}

