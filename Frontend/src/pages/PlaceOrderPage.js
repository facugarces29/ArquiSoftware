import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Review from '../components/Review';
import { useStateValue } from "../StateProvider";
import { actionTypes  } from "../reducer";

const useStyles = makeStyles((theme) => ({
  appBar: {
    position: 'relative',
  },
  layout: {
    width: 'auto',
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(2) * 2)]: {
      width: 600,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  },
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(3),
    padding: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
      marginTop: theme.spacing(6),
      marginBottom: theme.spacing(6),
      padding: theme.spacing(3),
    },
  },
  stepper: {
    padding: theme.spacing(3, 0, 5),
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
  },
  button: {
    marginTop: theme.spacing(3),
    marginLeft: theme.spacing(1),
  },
  center: {
    alignItems:'center',
    justifyContent:'center',
  },
}));

let couldCreateOrder = false;
async function postOrder( userID ){
  return await fetch('http://localhost:8080/order', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({"user_id":userID})
    })
    .then(response => {
      if (response.status === 400 || response.status === 401)
      {
        couldCreateOrder = false;
        return response.json();
      }
      couldCreateOrder = true;
      return response.json();
    })
}

let couldCreateDetail = false;
async function postDetail( orderID , product ){
  return await fetch('http://localhost:8080/order/detail', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({"order_id":orderID, "product_id":product.id, "price":product.price, "quantity":product.quantity})
    })
    .then(response => {
      if (response.status === 400 || response.status === 401)
      {
        couldCreateDetail = false;
        return response.json();
      }
      couldCreateDetail = true;
      return response.json();
    })
}

export default function Checkout() {
  const classes = useStyles();
  const [placed, setPlaced] = useState(false);
  const [error, setError] = useState(false);
  const [{ user, basket }, dispatch] = useStateValue();

  let orderID = 0;


  const placeOrder = (e) => {
    e.preventDefault();
    postOrder(user.userID).then(data => {
      if(couldCreateOrder === true){
        orderID = data.id;
        console.log("fuera",orderID);
        basket.map((item) => {
          console.log("item",item);
          console.log("dentro",orderID);
          postDetail(orderID , item).then(detailData => {
            if (couldCreateDetail === false){
              setError(true);
              return;
            }
          })
          return;
        })
      }
    })
    if(error == true){
      return;
    }else{
      setPlaced(true);
      dispatch({
        type: actionTypes.EMPTY_BASKET,
        basket: [],
      });
    }
  }

  return (
    <React.Fragment>
      <CssBaseline />
      <main className={classes.layout}>
        <Paper className={classes.paper}>
          <Typography component="h1" variant="h4" align="center">
            Checkout
          </Typography>
          <React.Fragment>
            { placed ? (
              <React.Fragment>
                <div >
                <Typography variant="h5" gutterBottom>
                  Thank you for your order.
                </Typography>
                <Typography variant="subtitle1">
                  Your order has been placed.
                </Typography>
                </div>
              </React.Fragment>
            ) : (
              <React.Fragment>
                <Review/>
                <div className={classes.buttons}>
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={placeOrder}
                    className={classes.button}
                  >
                    Place order
                  </Button>
                </div>
              </React.Fragment>
            )}
          </React.Fragment>
        </Paper>
      </main>
    </React.Fragment>
  );
}