import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Collapse from '@material-ui/core/Collapse';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import accounting from 'accounting'
import { AddShoppingCart } from '@material-ui/icons';

const useStyles = makeStyles((theme) => ({
    root: {
      maxWidth: 345,
    },
    action: {
      marginTop: "1rem",
    },
    media: {
      height: 0,
      paddingTop: "56.25%", // 16:9
    },
    expand: {
      transform: "rotate(0deg)",
      marginLeft: "auto",
      transition: theme.transitions.create("transform", {
        duration: theme.transitions.duration.shortest,
      }),
    },
    expandOpen: {
      transform: "rotate(180deg)",
    },
  }));
  

export default function Product({ productName , isInStock , productImage , productPrice , addToCart , description , productCategory }) {
  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  return (
    <Card className={classes.root}>
      <CardHeader
       action={
        <Typography
          className={classes.action}
          variant='h5'
          color='textSecondary'
        >
          {accounting.formatMoney(productPrice)}
        </Typography>
      }
        title={productName}
        subheader={ isInStock ? 'in stock' : 'no stock' }
      />
      <CardMedia 
        className={classes.media} 
        image={productImage} 
        title={productName} />
        <CardContent>
        <Typography variant='body2' color='textSecondary' component='p'>
          {productCategory}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton 
        aria-label='Add to Cart' 
        onClick={addToCart}>
          <AddShoppingCart 
            fontSize='large' />
        </IconButton>
        <IconButton
          className={clsx(classes.expand, {
            [classes.expandOpen]: expanded,
          })}
          onClick={handleExpandClick}
          aria-expanded={expanded}
          aria-label="show more"
        >
          <ExpandMoreIcon />
        </IconButton>
      </CardActions>
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <CardContent>
          <Typography paragraph>Description:</Typography>
          <Typography paragraph>{description}</Typography>
        </CardContent>
      </Collapse>
    </Card>
  );
}