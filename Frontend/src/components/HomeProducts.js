import { makeStyles } from "@material-ui/core/styles";
import { Grid, CssBaseline } from "@material-ui/core";
import Product from "../components/Product";
import { useEffect, useState } from "react";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
}));

const HomeProducts = () => {
  const classes = useStyles();
  
  const url = "http://127.0.0.1:8080/";
  
  const [products, setProducts] = useState();

  const fetchApi = async () => {
    const response = await fetch(url);
    const responseJSON = await response.json();
    setProducts(responseJSON);
  }

  useEffect(() => {
    fetchApi();
  }, []);
  
  return (
    <>
      <CssBaseline />
      <div className={classes.root}>
        <Grid container spacing={3}>
          { !products ? "cargando" : 
            products.map((product, index) => {
              
              let stock = true;
              if (product.stock === 0) {
                stock = false;
              }

              return  <Grid item xs={12} sm={6} md={4} lg={3}>
                        <Product
                          name={product.name}
                          price={product.price}
                          description={product.description}
                          image={product.image}
                          stock={product.stock}
                        />
                      </Grid>
            })

          }
          </Grid>
      </div>
    </>
  );
};

export default HomeProducts;