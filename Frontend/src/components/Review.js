import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Grid from '@material-ui/core/Grid';
import { useStateValue } from "../StateProvider";
import Box from '@material-ui/core/Box';
import { getBasketTotal } from "../reducer";
import accounting from "accounting";

const useStyles = makeStyles((theme) => ({
  listItem: {
    padding: theme.spacing(1, 0),
  },
  total: {
    fontWeight: 700,
  },
  title: {
    marginTop: theme.spacing(2),
  },
}));

export default function Review({address}) {
  const classes = useStyles();
  const [{ basket, user }, dispatch] = useStateValue();

  const payments = [
    { name: 'Card type', detail: 'Visa' },
    { name: 'Card holder', detail: user.username },
    { name: 'Card number', detail: 'xxxx-xxxx-xxxx-1234' },
    { name: 'Expiry date', detail: '04/2024' },
  ];


  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Order summary
      </Typography>
      <List disablePadding>
        {basket.map((item) => (
           <ListItem className={classes.listItem} key={item.name}>
           <ListItemText primary={item.name} secondary={`Quantity: ${item.quantity}`} />
           <Typography variant="body2">{accounting.formatMoney(item.price * item.quantity)}</Typography>
         </ListItem>
        ))}
        <ListItem className={classes.listItem}>
          <ListItemText primary="Total" />
          <Typography variant="subtitle1" className={classes.total}>
            <h5>{accounting.formatMoney(getBasketTotal(basket), "$")}</h5>
          </Typography>
        </ListItem>
      </List>
      <Grid container spacing={2}>
        <Grid item container direction="column" xs={12} sm={6}>
          <Typography variant="h6" gutterBottom className={classes.title}>
            Payment details
          </Typography>
          <Grid container>
            {payments.map((payment) => (
              <React.Fragment key={payment.name}>
                <Grid item xs={6}>
                  <Typography gutterBottom>{payment.name}</Typography>
                </Grid>
                <Grid item xs={6}>
                  <Typography gutterBottom>{payment.detail}</Typography>
                </Grid>
              </React.Fragment>
            ))}
            </Grid>
        </Grid>
      </Grid>
    </Box>
  );
}