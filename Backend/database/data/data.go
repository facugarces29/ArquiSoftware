package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	addressModel "proyecto/ArquiSoftware/model/address"
	orderModel "proyecto/ArquiSoftware/model/order"
	productModel "proyecto/ArquiSoftware/model/product"
	userModel "proyecto/ArquiSoftware/model/user"
)

func InsertData(db *gorm.DB) {
	// Insert data
	log.Info("Inserting data...")

	//Inserting users
	err := db.First(&userModel.User{}).Error

	if err != nil {
		db.Create(&userModel.User{Name: "lautaro", LastName: "Saenz", UserName: "lautarose", Email: "abcdefg@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Joaco", LastName: "Reyero", UserName: "joacoreyero", Email: "12345@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Facundo", LastName: "Garces", UserName: "Facuelcapo", Email: "asasas@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Hernan", LastName: "Lachampionliga", UserName: "hernanchampion", Email: "hernan@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Saul", LastName: "Hudson", UserName: "slash", Email: "slashGNR@gmail.com", Pwd: "hola123"})
	}

	//Inserting Categories

	err = db.First(&productModel.Category{}).Error

	if err != nil {
		db.Create(&productModel.Category{Name: "Electronics"})
		db.Create(&productModel.Category{Name: "Cars"})
		db.Create(&productModel.Category{Name: "Fitness"})
		db.Create(&productModel.Category{Name: "Toys"})
		db.Create(&productModel.Category{Name: "Construction"})
		db.Create(&productModel.Category{Name: "Supermarket"})
		db.Create(&productModel.Category{Name: "Babies"})
		db.Create(&productModel.Category{Name: "Fashion"})
	}

	//Inserting products

	err = db.First(&productModel.Product{}).Error

	if err != nil {
		//Electronics
		db.Create(&productModel.Product{CategoryID: 1, Name: "Asus Laptop", Description: "Personal laptop, intel i5, 8gb ram, fullhd screen, windows 10", Price: 100000, Stock: 10, Image: "https://dlcdnwebimgs.asus.com/gain/9a249989-9272-4645-9a0a-37e29000e369/w800"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Logitech Mouse", Description: "Gaming Mouse 16000dpi", Price: 8000, Stock: 15, Image: "https://resource.logitechg.com/w_1000,c_limit,q_auto,f_auto,dpr_auto/d_transparent.gif/content/dam/products/gaming/gaming-mice/g305-lightspeed-wireless-gaming-mouse/g304-g305-lightspeed-wireless-gaming-mouse21.png?v=1"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Hyperx Keyboard", Description: "Keyboard rgb, cherry, us distribution", Price: 12000, Stock: 4, Image: "https://media.kingston.com/hyperx/product/hx-product-keyboard-alloy-fps-rgb-hxkb1ss2uk-1-sm.jpg"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Samsung Monitor", Description: "Resolution: 1920x1080, 144hz, 32 inches", Price: 125000, Stock: 2, Image: "https://images.samsung.com/is/image/samsung/latin-t35f-lf24t350fhlxzp-frontblack-301443711?$650_519_PNG$"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "RX 6700XT", Description: "AMD es un fabricante estadounidense de placas de video, por su tecnología se ha destacado en crear procesadores de alta gama que permiten un excelente funcionamiento del motor gráfico de tu computadora", Price: 189000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_655650-MLA46792682690_072021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "MSI RTX 3060", Description: "Nvidia es el fabricante líder de placas de video; su calidad asegura una experiencia positiva en el desarrollo del motor gráfico de tu computadora", Price: 114000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_667938-MLA48657941510_122021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Intel Celeron G5905 BX80701G5905", Description: "Productividad y entretenimiento, todo disponible en tu computadora de escritorio", Price: 6700, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_677873-MLA44347499239_122020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Intel Pentium G6400 BX80701G6400", Description: "Núcleos: el corazón del procesador", Price: 10000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_768452-MLA43682182878_102020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "MD Athlon 3000G", Description: "Al estar desbloqueado, podrás realizar overclocking", Price: 18000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_620009-MLA41419343230_042020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "AMD Ryzen 7 5700G 100-100000263BOX de 8 núcleos y 4.6GHz de frecuencia con gráfica", Description: "Clave en el rendimiento de tu computadora de escritorio", Price: 50000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_940934-MLU47593217192_092021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "i5-10400F", Description: "Productividad y entretenimiento", Price: 23000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_602530-MLA43003993713_082020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Memoria RAM Fury Beast DDR4", Description: "Con su tecnología DDR4, mejorará el desempeño de tu equipo", Price: 13000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_859341-MLA49041755528_022022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Fury Beast DDR4 RGB color negro 8GB", Description: "", Price: 5700, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_908770-MLA48636149604_122021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Samsung F22T35 22", Description: "Un monitor a tu medida", Price: 37000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_796587-MLA46165231779_052021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Samsung Odyssey G3 F24G35T 24 ", Description: "Una experiencia visual de calidad", Price: 120000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_926028-MLA48705312123_122021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Samsung F27T350FHL 27", Description: "on tu pantalla LED no solo ahorrás energía, ya que su consumo es bajo, sino que vas a ver colores nítidos y definidos en tus películas o series favoritas.", Price: 110000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_837353-MLA45884040838_052021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Razer Hammerhead True Wireless X", Description: "Experimentá la adrenalina de sumergirte en la escena de otra manera", Price: 7800, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_640507-MLA49602463222_042022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "JBL Wave 100TWS JBLW100TWS", Description: "El formato perfecto para vos", Price: 8900, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_875020-MLA47919581799_102021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Noga Aris NG-BT469", Description: "El diseño over-ear brinda una comodidad insuperable gracias a sus suaves almohadillas", Price: 2000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_962972-MLA42763657387_072020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Genérica F9-5", Description: "Bluetooth de última generación", Price: 1600, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_938092-MLA45480677826_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Amazon Echo Dot 3rd Gen", Description: "Echo Dot 3rd Gen es el parlante más popular de Amazon", Price: 10000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_917766-MLA43059544028_082020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Amazon Echo Dot 4th Gen with clock", Description: "Sumate al mundo de la inteligencia artificial con el Amazon Echo Dot 4th Gen with clock que ofrece soluciones para automatizar algunas tareas en tu hogar.", Price: 100000, Stock: 10, Image: ""})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Amazon Echo Show 8", Description: "Alexa te acompaña", Price: 19000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_760260-MLA49971550922_052022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Lenovo Smart Cloc", Description: "Sistema de audio multi-room", Price: 17000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_759861-MLA46270225231_062021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Edifier 2.0 R1580MB", Description: "Versátil y funciona", Price: 26000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_626377-MLA41638649339_052020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Thonet & Vander Vertrag BT", Description: "Thonet & Vander Vertrag BT ofrece un sonido natural", Price: 36000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_821841-MLA46257376523_062021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Harman Kardon Aura Studio 3", Description: "Gran potencia", Price: 63000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_865586-MLA49374644391_032022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "JBL Boombox 2", Description: "Olvidate del amplificado", Price: 100000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_893530-MLA45178759714_032021-O.webp"})

		//Fitness
		db.Create(&productModel.Product{CategoryID: 3, Name: "Randers Multigym", Description: "75kg, more than 30 exercises, aluminium", Price: 85000, Stock: 3, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Smart Tech Fixed Bycicle", Description: "Spinning exercises, high endurance", Price: 35000, Stock: 1, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Pull-Up Bar", Description: "Pull-up bar", Price: 9500, Stock: 6, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Dumbell and Kettlebell kit", Description: "2x10kg dumbell,2x5kg kettlebell", Price: 7000, Stock: 5, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Suplemento en polvo", Description: "Star Nutrition Creatine Monohydrate creatina en pote de 300g", Price: 6498, Stock: 78, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_705680-MLA48462687556_122021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Piso Encastrable", Description: "Pack X 4 Goma Eva 60x60cm Entrenamiento Gym", Price: 2689, Stock: 15, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_967182-MLA45740888976_042021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Banco De Pesas", Description: "Pecho Plano Inclinado Y Piernas Envío Gratis", Price: 28000, Stock: 8, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_651443-MLA32945071209_112019-F.webp"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Kit Gym Mancuerna", Description: "Barra Topes Rosca Fitness + Discos 30 Kg", Price: 12699, Stock: 35, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_915514-MLA46478611709_062021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Multigimnasio", Description: "75kg C/estruct. P/lingotera World Fitness 7116", Price: 85000, Stock: 55, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_624524-MLA31047194237_062019-F.webp	"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Banco De Pesas Pecho Plano", Description: "Banco Para Pecho Plano E Inclinado JDM SPORTS", Price: 28000, Stock: 3, Image: "https://http2.mlstatic.com/D_NQ_NP_651443-MLA32945071209_112019-O.webp%22%7D"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "BARRA RECTA MACIZA", Description: " Pesa sola aproximadamente 8 kg", Price: 15000, Stock: 3, Image: "https://http2.mlstatic.com/D_NQ_NP_684565-MLA48438349691_122021-O.webp%22%7D"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "TOBILLERAS DE 3 KG SUPER REFORZADAS", Description: "Con ajuste regulable", Price: 1300, Stock: 120, Image: "https://http2.mlstatic.com/D_NQ_NP_718250-MLA44050687382_112020-O.webp%22%7D"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Kit Set Entrenamiento Funcional Fitness", Description: "Cross Funcional Regalo Deportivo", Price: 85000, Stock: 3, Image: "https://http2.mlstatic.com/D_NQ_NP_816975-MLA46738837354_072021-O.webp%22%7D"})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Pesa Rusa O Kettlebell 8 Kg Pvc Funcional Crossfit", Description: "INDKET", Price: 1800, Stock: 13, Image: "https://http2.mlstatic.com/D_NQ_NP_808866-MLA44549965327_012021-O.webp%22%7D"})

		//Toys
		db.Create(&productModel.Product{CategoryID: 4, Name: "Skateboard", Description: "Skateboard", Price: 10000, Stock: 3, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Hot-Wheels Camaro", Description: "1980 Camaro", Price: 2000, Stock: 10, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Nerf n-strike", Description: "nerf n-strike", Price: 5500, Stock: 2, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Lego Spidey", Description: "100 pcs", Price: 7990, Stock: 1, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Juego de mesa ", Description: "Monopoly Popular Hasbro 840", Price: 5690, Stock: 13, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_708184-MLA44676742923_012021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Pileta inflable", Description: "Redonda Bestway 51026 de 1520mm x 300mm 282L azul", Price: 3799, Stock: 56, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_645105-MLA44938932100_022021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Juego de Cartas", Description: "Uno Ruibal", Price: 1500, Stock: 11, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_942888-MLA46738841701_072021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Ametralladora", Description: "Con Sonido Disparo A Friccion Juguete Niños", Price: 606, Stock: 78, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_609493-MLA48088057407_112021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Pistola Rifle", Description: "Con Mira X-shot Hawk Eye - 12 Dardos. 01186", Price: 5040, Stock: 75, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_739521-MLA44945098920_022021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Peluche", Description: "Mascota Sorpresa Furball Scruff A Luvs Babies", Price: 2121, Stock: 35, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_605322-MLA49333794356_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Peluche", Description: "Huggy Wuggy 48cms Velcro Poppy Playtime", Price: 1799, Stock: 75, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_653489-MLA49526402611_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Peluches", Description: "Pokemon Go 20 Cm Varios Modelos Pikachu Pokemones", Price: 1289, Stock: 89, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_668216-MLA48876225229_012022-F.webp"})

		//Cars
		db.Create(&productModel.Product{CategoryID: 2, Name: "Motor Completo Chevrolet Cruze 1.4 T 2017", Description: "Somos una empresa dedicada a la venta de autopartes, por Mayor y Menor, con mas de treinta añosos de experiencia en el mercado", Price: 240000, Stock: 30, Image: "https://http2.mlstatic.com/D_NQ_NP_744704-MLA49772787150_042022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Volante Deportivo Collino Tapizado Gamuza Veloce R W 355 Mm", Description: "C-FORCE - VOLANTES - TRACK", Price: 20000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_789721-MLA45533806368_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Selectora Rapida Collino Caja Lancia Con Anti Primera Fiat", Description: "SELECTORA LANCIA ANTI 1 FIAT UNO - FIAT 128", Price: 49000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_620982-MLA45564839396_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Tren Delantero Ford Falcon Linea Nueva", Description: "SE INSTALA SIN REFORMAS DE NINGÚN TIPO EN LOS ANCLAJES NI SOLDADURA", Price: 400000, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_630938-MLA45580404286_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Bujes Elastomero Elasticos", Description: "BUJES ELASTICOS FALCON", Price: 24000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_810657-MLA45496393449_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Difusores Collino Carburador Dino 36-36", Description: "DIFUSORES DINO 36/36", Price: 40000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_735944-MLA47585401318_092021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Filtro De Aire Torino Para Carburador Dino Caresa", Description: "FILTRO AIRE CHEVROLET", Price: 18700, Stock: 21, Image: "https://http2.mlstatic.com/D_NQ_NP_849325-MLA47262398800_082021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Base Adaptadora Collino Holley 4040", Description: "BASE ADAPTADORA DE ALUMINIO PARA CARBURADORES", Price: 10000, Stock: 12, Image: "https://http2.mlstatic.com/D_NQ_NP_824211-MLA46980709218_082021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Recuperador De Aceite Aluminio", Description: "CATCH CAN 450 CM3 SILVER SALIDA DERECHA", Price: 12000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_911442-MLA45626997482_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Pedal Acelerador Collino Go!", Description: "PEDAL ACELERADOR GO", Price: 19000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_811025-MLA45549635781_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Freno De Mano Vertical", Description: "USO 1/4 DE MILLA", Price: 23000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_810159-MLA48992770698_022022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Bomba De Freno De Mano Competicion", Description: "BOMBA DE FRENO DE MANO", Price: 19000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_982433-MLA45548672525_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Corrector De Levas", Description: "FASTER - FASTER VW - FASTER POLEAS VW", Price: 12000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_929834-MLA47007684813_082021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Corrector De Levas Vw Gol Ap Diente Redond + Distrib", Description: "DISTIRIBUCION VW AP 8V", Price: 40000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_920808-MLA50207550530_062022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Tensor Respaldo Correa Distribucion", Description: "ENSOR RESPALDO AP", Price: 4000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_679769-MLA45566651940_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Tapa De Valvulas Aluminio", Description: "TAPA VW AP LISA", Price: 16000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_885042-MLA45598496216_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Oring Sello Para Tapa De Valvulas", Description: "ORING TAPA FIAT", Price: 4000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_815025-MLA47660220422_092021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Refuerzo Inferior Collino Fiat 128 S.e. Sin Tensores", Description: "REFUERZO INF 128 S.E.", Price: 17000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_611784-MLA47009976636_082021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Strut Bar Collino Fiat 128", Description: "STRUT BAR FIAT 128", Price: 10000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_782810-MLA47544981290_092021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Strut Bar Collino Fiat 1 Uno Roja Barra Torretas Rotulada", Description: "PESO TOTAL: 1450 GRS", Price: 20000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_754814-MLA44216533000_122020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Porta Rotulas Extremos Para Tensores Fiat 1", Description: "DESPIECE COLLINO - SUSPENSION - TREN DELANTERO FIAT", Price: 10000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_958651-MLA45629018688_042021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Varillas Rotula Selectora Caja Lancia", Description: "VARILLAS ROTULAD LANCIA", Price: 14000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_732993-MLA47248089092_082021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Polea Cigueñal Fiat 1 Motor Tipo Faste", Description: "POLEA EN V FIAT TIPO", Price: 4000, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_875911-MLA49715175102_042022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Engranaje De Bronce", Description: "ENGRANAJE DISTR. F TIPO", Price: 8400, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_795695-MLA47916480707_102021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Engranaje D Bronce Collino Distribuidor Motor Fiat 600 S", Description: "ENGRANAJE DISTR. F 600", Price: 8300, Stock: 20, Image: "https://http2.mlstatic.com/D_NQ_NP_869997-MLA47583681702_092021-O.webp"})

		//Construction
		db.Create(&productModel.Product{CategoryID: 5, Name: "Bomba De Agua Presurizadora", Description: "BOMBA PRESURIZADORA 8,5M 220V 30L/MIN LUSQTOFFLPS15-8.5Z", Price: 7305, Stock: 1, Image: "https://http2.mlstatic.com/D_NQ_NP_897453-MLA49312640274_032022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Ducha Cuadrada", Description: "Acero Inoxidable 25X25 cm + barral de 42 cm hasta la curva, distancia total 50 cm", Price: 3824, Stock: 1, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_970470-MLA31384529250_072019-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Pintura Latex Interior", Description: "Albalatex Blanco Mate X 20 Lts - Kromacolor", Price: 15945, Stock: 1, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_763816-MLA40746058396_022020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Griferia Ducha", Description: "Peirano 80-010 Lorca Embutir Con Transferenci", Price: 7955, Stock: 1, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_880688-MLA48919942667_012022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Carro Zorra Carretilla", Description: "Portabultos Ropal Capacidad 200 Kg", Price: 8074, Stock: 2, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_756180-MLA45212431204_032021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Cerradura", Description: "Prive 200 Reforzada Herrero Casa Puerta Oferta", Price: 1305, Stock: 1, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_657604-MLA31099034696_062019-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Tejido Metal Electrosoldado", Description: "BMalla Alambre 25x25mm 50cm Alto", Price: 395, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_970963-MLA46245939770_062021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Greiferia", Description: "Greiferia", Price: 19999, Stock: 12, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_936614-MLA49293925811_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Cable", Description: "Cable Subterraneo 2x6 Mm X 50 Mts / T", Price: 12000, Stock: 21, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_792191-MLA49523940998_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Kit Solar", Description: "Kit Solar Completo Must Autoinstalable 1000w Panel Bateria", Price: 49795, Stock: 15, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_663230-MLA49476173586_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Inodoro Con Mochila", Description: " Ferrum Largo Bidet Juego De Baño Marina ", Price: 176705, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_878233-MLA42044156380_062020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Protector Tension", Description: "Baw Trifasico Monofasico 63a 220v 380v", Price: 12425, Stock: 43, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_881380-MLA49814760093_042022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Pinza Amperimetrica", Description: "Digital BAW MT87 400A", Price: 966, Stock: 14, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_886517-MLA46819474452_072021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Llave De Luz", Description: "Armada Sica Life 2 Tomas Binorma | Pack X 10", Price: 2285, Stock: 9, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_772153-MLA47576802787_092021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Zapatilla", Description: "Prolongador Kalop Multiple 6 Tomas +1,5 Mts Cable", Price: 1556, Stock: 94, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_862649-MLA28450471735_102018-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Placas 3d", Description: "Autoadhesivas Ladrillos Tipo Pared Anti Humedad", Price: 886, Stock: 104, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_905257-MLA49184838843_022022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Porcelanato", Description: "Beige Pulido 60x60 Primera Rectificado", Price: 3300, Stock: 89, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_996117-MLA41603221030_052020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 5, Name: "Grifo de cocina", Description: "Monocomando Libercam Rm 300 plateado acabado cromado", Price: 2700, Stock: 45, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_792047-MLA47576602069_092021-F.webp"})

		//SuperMarket
		db.Create(&productModel.Product{CategoryID: 6, Name: "Lavandina ", Description: "Ayudín Antisplash No Salpica 2l Limpieza", Price: 229, Stock: 13, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_704321-MLA48915290743_012022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Aceite De Oliva", Description: "Extra Virgen Nucete En Bidon 2 Litros", Price: 2026, Stock: 34, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_708716-MLA49169156071_022022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Vino Nicasia", Description: "Red Blend Malbec 750ml Caja X 6u", Price: 6, Stock: 6534, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_741907-MLA49421442983_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Kit Elvive Óleo Coco", Description: "Kit Elvive Óleo Coco", Price: 1726, Stock: 18, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_666135-MLA48600079323_122021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Cerveza", Description: "Pampa Ipa 473cc ", Price: 182, Stock: 1122, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_665310-MLA48598736996_122021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Cepillo Dental", Description: "Colgate Proplanet Sustentable Mango Aluminio", Price: 1586, Stock: 14, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_700646-MLA46764522049_072021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Whisky", Description: "Jack Daniel's Old No. 7 750 mL", Price: 5226, Stock: 4, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_712680-MLA44850207740_022021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Pañales", Description: "Pañales para adultos descartables Plenitud Ropa Interior Protect Plus G/XG x 16 u", Price: 2842, Stock: 75, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_908303-MLA43078780884_082020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Alimento para gatos", Description: "Agility Premium Urinary para gato adulto sabor mix en bolsa de 10 kg", Price: 5499, Stock: 13, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_758004-MLA50131748668_052022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 6, Name: "Alimento para perro", Description: " Pedigree Óptima Digestión Etapa 2 para perro adulto todos los tamaños sabor carne, pollo y cereales en bolsa de 21 kg", Price: 5126, Stock: 78, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_831071-MLA49900705686_052022-F.webp"})

		//Babies
		db.Create(&productModel.Product{CategoryID: 7, Name: "Cochecito Bebe ", Description: "Avanti Reversible Convertible En Moises Smile", Price: 21229, Stock: 10, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_600452-MLA48247784080_112021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Butaca ", Description: "infantil para auto Mega Baby Silverstone gris", Price: 13500, Stock: 8, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_679143-MLA42229530141_062020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Mochila ", Description: "Bolso Mama Para Mamaderas Bebe Termica Garantia", Price: 4999, Stock: 6, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_770589-MLA49364224431_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Silla", Description: "Para comer 3 Alturas Reclinable Verde - Wondrus", Price: 6599, Stock: 14, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_760440-MLA49262538432_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Leche", Description: "Leche de fórmula en polvo sin TACC Nutricia Bagó Nutrilon Profutura 3 en lata de 800g por 6 unidades - 12 meses 2 años", Price: 9420, Stock: 89, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_807227-MLA42346942112_062020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Pelela", Description: "Educador Inodoro Avanti Con Tapa Y Mochila Reductor", Price: 5389, Stock: 28, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_732663-MLA44040262961_112020-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Chupetes", Description: "Bibs Pacifier Set X2 Unds Libre Bpa Latex Natural", Price: 2099, Stock: 189, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_955630-MLA46793204659_072021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Practicuna", Description: "Carestino Venecia Aguamarina Completa + Bolso", Price: 20999, Stock: 9, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_622201-MLA50111810019_052022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Frazada", Description: "Para Bebe Con Corderito Suave Abrigada Calentita Cuna", Price: 2300, Stock: 54, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_608158-MLA46207906096_052021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 7, Name: "Faja Anticólicos", Description: "Aroterm Con Almohadilla De Semillas Lavanda", Price: 1229, Stock: 13, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_708947-MLA46164633217_052021-F.webp"})

		//Fasion
		db.Create(&productModel.Product{CategoryID: 8, Name: "Zapatillas", Description: "Jaguar Oficial Deportiva Art. # 918 Urbana Hombre Y Mujer", Price: 3430, Stock: 38, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_968157-MLA27429372161_052018-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Jogging", Description: "Hombre Chupin Deportivo Babucha Pantalon No Jean", Price: 1700, Stock: 27, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_697340-MLA45393217065_032021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Campera", Description: "Inflable Negra Super Abrigada Con Capucha", Price: 7899, Stock: 21, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_664776-MLA49302750254_032022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Polera", Description: "De Morley. Clasica, Elastizada, Mas De 15 Colores!!!", Price: 1389, Stock: 31, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_675818-MLA46172311941_052021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Zapartillas", Description: "Coreracer adidas Sport 78 Tienda Oficial", Price: 11999, Stock: 30, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_828977-MLA50128052624_052022-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Buzo", Description: "Canguro Capucha Hombre Slim Fit Hoodie Camperas", Price: 3729, Stock: 123, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_961834-MLA31615108584_072019-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Carolina Herrera CH Men EDT 100 ml para hombre", Description: "Durante las últimas décadas, Carolina Herrera se convirtió en una de las diseñadoras más prestigiosas del mundo. Sus perfumes marcan tendencia y se convierten en el accesorio invisible por excelencia", Price: 12000, Stock: 300, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_829567-MLA45187812495_032021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Versace Eros", Description: "La fragancia Eros eau de toilette ofrece la frescura perfecta para épocas de calor. Al caracterizarse por sus aromas ligeros podés usar la cantidad que quieras sin miedo a excederte.", Price: 14000, Stock: 146, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_678443-MLA46088855665_052021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Paco rabanne", Description: "Respaldado por décadas de memorables fragancias que llevan su nombre, Paco Rabanne afirma una identidad distintiva resultante de la síntesis del diseño contemporáneo de vanguardia y la artesanía radical.", Price: 13000, Stock: 146, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_766773-MLA45925248250_052021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Antonio bandera King seduction", Description: "La fragancia King of Seduction eau de toilette ofrece la frescura perfecta para épocas de calor. Al caracterizarse por sus aromas ligeros podés usar la cantidad que quieras sin miedo a excederte.", Price: 4000, Stock: 1467, Image: "https://http2.mlstatic.com/D_NQ_NP_748501-MLA47574177150_092021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Paco Rabanne 1 million", Description: "Envuelto en frasco con forma de lingote, 1 Million es el perfume para el hombre que ve la vida de oro, para el que ningún sueño es demasiado grande. Su aroma es cautivador, de esos que dejan huella en cada lugar.", Price: 13000, Stock: 100, Image: "https://http2.mlstatic.com/D_NQ_NP_955869-MLA44750390454_012021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Afeitadora Philips oneblade", Description: "Innovación y calidad son los máximos valores que representan a Philips. Desde sus inicios, esta marca de origen neerlandesa se ha propuesto ser una referente mundial en el mercado de tecnologías.", Price: 4000, Stock: 268, Image: "https://http2.mlstatic.com/D_NQ_NP_976065-MLA43279963882_082020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Tratamiento Antiedad Con Ácido Hialuronico Hyalu B5 Serum", Description: "Hyalu B5 Serum de La Roche Posay es un tratamiento antiedad para pieles sensibles compuesto por ácido hialurónico de doble peso molecular, vitamina B5 para disminuir la inflamación y agua termal.", Price: 8300, Stock: 268, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_930033-MLA45107695099_032021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Set Antiedad Regina Serum Vitamina C & Hialuronico", Description: "Set Antiedad de Hidratación Profunda Regina", Price: 7400, Stock: 268, Image: "https://http2.mlstatic.com/D_NQ_NP_902546-MLA41851425858_052020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Rubor Tonalizador En Sombra Cremosa Para Contour Regina", Description: "Fácil de aplicar, textura ligera, resistente a la transpiración, se funde sin alterar el maquillaje para broncear, definir o contornear.", Price: 1900, Stock: 300, Image: "https://http2.mlstatic.com/D_NQ_NP_932270-MLA49794717090_042022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Máscara de pestañas Regina Reina 10g color negro", Description: "Destacá tu mirada con esta máscara Regina que define y cuida tus pestañas", Price: 2300, Stock: 300, Image: "https://http2.mlstatic.com/D_NQ_NP_884645-MLA42258006174_062020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Set De Iluminadores Regina", Description: "Set de iluminadores Regina", Price: 5700, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_723771-MLA47678026753_092021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Regina Corrector De Ojeras Sin Arrugas Inicio", Description: "INICIO® Corrector de Ojeras Sin arrugas Regina", Price: 2999, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_986258-MLA44878941421_022021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Lápiz Labial Mate Regina", Description: "Lápiz Labial Mate Regina 1.20g", Price: 2700, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_991181-MLA50126319988_052022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Locion Tonica | Hidra Tonic", Description: "Loción Tonica | Hidra Tonic Regina Loción refrescante y suavizante.", Price: 1840, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_723439-MLA42094949044_062020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Labial Liquido Mate Regina", Description: "Labial Líquido Mate Regina 3.5 ml Larga Duración #CrueltyFree", Price: 1650, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_700031-MLA41815586481_052020-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "La Roche Posay Retinol B3 Serum X 30m", Description: "Nuevo Serum Antiedad Retinol B3 es un tratamiento antiedad que trata arrugas profundas y tono irregular apto para pieles sensibles.", Price: 11850, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_937976-MLA45107815878_032021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Hidratante Intensivo Ligero La Roche Posay", Description: "Hydraphase HA ligera 50ml RENOVACION", Price: 6600, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_975102-MLA46219790556_052021-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "The Ordinary Acido Hialuronico 2% + B5 Original Serum Facial", Description: "Es un serúm con ácido hialurónico vegano totalmente puro", Price: 3850, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_764873-MLA48790053945_012022-O.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Vichy Liftactiv Supreme H.a. Epidermic Filler Serum Antiedad", Description: "Recarga el 100% de la pérdida diaria de Ácido Hialurónico en la piel*. Liftactiv H.A", Price: 8800, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_2X_911922-MLA45130616239_032021-F.webp"})
		db.Create(&productModel.Product{CategoryID: 8, Name: "Vichy Normaderm Phytosolution Gel Purificante Concentrado 40", Description: "Phytosolution - Gel Purificante concentrado", Price: 7200, Stock: 278, Image: "https://http2.mlstatic.com/D_NQ_NP_960296-MLA49771551680_042022-O.webp"})

	}

	//Inserting addresses

	err = db.First(&addressModel.Address{}).Error

	if err != nil {
		db.Create(&addressModel.Address{UserID: 1, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Eugenio Garzon 467"})
		db.Create(&addressModel.Address{UserID: 2, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Ituzaingo 200"})
		db.Create(&addressModel.Address{UserID: 3, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "General Paz 344"})
		db.Create(&addressModel.Address{UserID: 4, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Velez Sarsfield 1500"})
		db.Create(&addressModel.Address{UserID: 5, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Belgrano 110"})
	}

	//Inserting orders

	err = db.First(&orderModel.Order{}).Error

	if err != nil {
		db.Create(&orderModel.Order{UserID: 1})
		db.Create(&orderModel.OrderDetail{OrderID: 1, ProductID: 15, Price: 1000, Quantity: 2})
		db.Create(&orderModel.OrderDetail{OrderID: 1, ProductID: 23, Price: 1500, Quantity: 1})
		db.Create(&orderModel.Order{UserID: 1})
		db.Create(&orderModel.OrderDetail{OrderID: 2, ProductID: 56, Price: 300, Quantity: 3})
		db.Create(&orderModel.Order{UserID: 2})
		db.Create(&orderModel.OrderDetail{OrderID: 3, ProductID: 105, Price: 2842, Quantity: 2})
		db.Create(&orderModel.Order{UserID: 3})
		db.Create(&orderModel.OrderDetail{OrderID: 4, ProductID: 105, Price: 2842, Quantity: 2})
	}

	//manage errors...

	log.Info("Data inserted")
}
