import { makeStyles } from "@material-ui/core/styles";
import { Grid, CssBaseline } from "@material-ui/core";
import { useState , useEffect} from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Link from "@material-ui/core/Link";
import Box from "@material-ui/core/Box";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import Typography from "@material-ui/core/Typography";
import Container from "@material-ui/core/Container";
import { actionTypes } from "../reducer";
import { useStateValue } from "../StateProvider";
import { useNavigate } from "react-router-dom";
import OrderCard from "../components/OrderCard"
import { ContactSupportOutlined } from "@material-ui/icons";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
  card: {
    display:'flex',
    alignItems:'center',
    justifyContent:'center',
  },
  text:{
    display:'flex',
    alignItems:'center',
    justifyContent:'center',
  },
}));

const Orders = () => {
  const classes = useStyles();
  const [orders, setOrders] = useState();
  const [products, setProducts] = useState();
  const [{ user }, dispatch] = useStateValue();
  let idUser = 0;

  user? idUser = user.id : idUser = 0;

  console.log(idUser);
  
  const urlOrders = "http://127.0.0.1:8080/order/" + idUser;

  const urlProducts = "http://127.0.0.1:8080/product/";
  
  const fetchOrders = async () => {
    const response = await fetch(urlOrders);
    const responseJSON = await response.json();
    setOrders(responseJSON);
  }

  useEffect(() => {
    fetchOrders();
  }, []);

  var details = [];

  console.log(details);

  !orders ? console.log("no hay ordenes") : orders.map((order) => {
    details.push(order.order_details);
    console.log("order:" , order);
  })

  console.log("details: " , details);

  const fetchProducts = async ()=> {
    const response = await fetch(urlProducts);
    const responseJSON = await response.json();
    return responseJSON;
  }
  
  useEffect(() => {
    ! details? console.log("no hay detalles") : details.forEach(function(detail) {
      console.log("detail dentro use effect" , detail);
      fetchProducts(detail[0].product_id);
    })
  }, []);

  console.log("products:" , products);

  return (
    <>
      <CssBaseline/>
      <div className={classes.root}>
        <Grid container spacing={3} className={classes.root}>
          { !orders ? <Typography  variant='h3' color='textPrimary' className={classes.text}> todavia no has realizado ninguna orden.</Typography> : 
            orders.map((order) => {
            return  <Grid item xs={12} sm={12} md={12} lg={12} className={classes.card}>
                      <OrderCard
                        date={order.date}
                        amount={order.amount}
                      ></OrderCard>
                      </Grid>
            })

          }
          </Grid>
      </div>
    </>
  );
};

export default Orders;