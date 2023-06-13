import React, { useState } from 'react';
import Popup from 'reactjs-popup';
import ContentModal from './contentModal';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import CardActions from '@mui/material/CardActions';
import IconButton from '@mui/material/IconButton';
import StarBorderIcon from '@mui/icons-material/StarBorder';
import RadioButtonUncheckedIcon from '@mui/icons-material/RadioButtonUnchecked';

type Props = {
  article: Article;
  articleSourceTitle: string | undefined;
  isAdmin: boolean;
};

const ArticleCard: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState(false);
  const handleContentModalClose = () => {
    setIsContentModalOpen(false);
  };
  // get first image

  // add _blank to <a></a>
  return (
    <>
      <Card
        sx={{ maxWidth: 345 }}
        className="articleCard"
        onClick={() => setIsContentModalOpen(!isContentModalOpen)}
      >
        <CardMedia
          component="img"
          height="194"
          image="/static/images/cards/paella.jpg"
          alt="Paella dish"
        />
        <CardContent>
          <Typography className='articleCard__title' gutterBottom variant="h5" component="div">
            {props.article.title}
          </Typography>
          <Typography className='articleCard__source' variant="body2" color="text.secondary">
            {props.articleSourceTitle}
          </Typography>
        </CardContent>
        <CardActions disableSpacing className={props.isAdmin ? "d-none" : "d-block"}>
          <IconButton aria-label="read later">
            <StarBorderIcon />
          </IconButton>
          <IconButton aria-label="status">
            <RadioButtonUncheckedIcon />
          </IconButton>
        </CardActions>
      </Card>
      <Popup modal open={isContentModalOpen} onClose={handleContentModalClose}>
        <ContentModal
          article={props.article}
          handleContentModalClose={handleContentModalClose}
        />
      </Popup>
    </>
  );
};

export default ArticleCard;
