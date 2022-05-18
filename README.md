# Api backend e-commerce para Arquitectura de software en golang.

# REQUERIMIENTOS:

- Autentificacion de usuarios: login, permiso.
- Home app.
- Search/paginacion (buscar productos por palabras claves).
- Detalle del producto (Posibilidad de agregar productos al carrito 1 o mas unidades del mismo).
- Mi carrito (pagina de resumen con todos los productos cargados en mi carrito).
- Confirmacion de la compra (desde la pagina de resumen), resumen de la compra y direccion de la entrega (se asume que el pago fue aprobado).
- Pagina "mis compras" (listas de todas las ordenes que se hicieron con el usuario que esta logueado).

# Extra points:

- Paginacion (posibilidad de listar los productos por categoria).
- Pagina de detalle de las compras realizadas con los productos de la misma.

# Restful:

- /home                         GET
- /search/:id                   GET
- /login                        GET
- /cart                         GET
- /cart/buy                     PUT
- /buy                          PUT
- /products                     GET
- /products/:name               GET
- /products/categories          GET
- /products/categories/:name    GET
- /order                        GET
- /order/detail                 GET
- /purchases                    GET

