import { makeStyles } from "@material-ui/core/styles";
import { Grid, CssBaseline } from "@material-ui/core";
import Product from "../components/ProductCard";
import { Typography } from "@material-ui/core";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";


const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
  notFound: {
    width:'100%',
    justifyContent:'center',
    alignItems:'center',
  },
}));

const SearchPage = () => {
  const classes = useStyles();
  const [products, setProducts] = useState();
  const param = useParams();
  const url = `http://localhost:8080/search/${param.id}`;

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
          { !products ? <Grid item xs={12} sm={12} md={12} lg={12}> <Typography align='center' gutterBottom variant='h4'>No results related to your search </Typography> </Grid> : 
            products.map((product, index) => {
              let stock = true;
              if (product.stock === 0) {
                stock = false;
              }
              return  <Grid item xs={12} sm={6} md={4} lg={3}>
                        <Product
                          id = {product.id}
                          name={product.name}
                          price={product.price}
                          description={product.description}
                          image={product.image}
                          stock={product.stock}
                          category={product.category}
                        />
                      </Grid>
            })
          }
          </Grid>
      </div>
    </>
  );
};

export default SearchPage;