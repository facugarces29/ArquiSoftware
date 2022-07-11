import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export default function OrderCard({ image, amount, date , products , address }) {
  const classes = useStyles();

  const renderProducts = (productName, productPrice, productQuantity) => {
    return () => {
      <Typography variant="body2" color="textSecondary" component="p">
      -{productName}: price: ${productPrice}, quantity: {productQuantity}
      </Typography>
    }
  }

  return (
    <Card className={classes.root}>
      <CardActionArea>
        <CardMedia
            component="img"
            alt={products[0].name}
            height="140"
            image={image}
            title="product image"
          />
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {date}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            <strong>Products:</strong><br></br>
            {products.map((product) => (
              <div>
                <Typography  variant="body2" color="textSecondary" component="p">
                  - {product.name} <br></br> price: ${product.price} <br></br> quantity: {product.quantity} <br></br>
                </Typography>
              </div>
          ))}
          <strong>- Total amount: </strong>${amount}<br></br>
          <strong>- Address: </strong> {address.address_line}, {address.zip}, {address.city}, {address.state} 
          
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}