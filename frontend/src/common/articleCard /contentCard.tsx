import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import CardActions from '@mui/material/CardActions';
import IconButton from '@mui/material/IconButton';
import StarBorderIcon from '@mui/icons-material/StarBorder';
import RadioButtonUncheckedIcon from '@mui/icons-material/RadioButtonUnchecked';
import { CARD_MAX_WIDTH } from '.';

type Props = {
  articleTitle: string;
  articleSourceTitle: string | undefined;
  isAdmin: boolean;
  handleModal: () => void;
  content: string;
};

const ContentCard: React.FC<Props> = (props: Props) => {
  return (
    <Card
      sx={{ maxWidth: CARD_MAX_WIDTH }}
      className="articleCard"
      onClick={props.handleModal}
    >
      <CardContent>
        <Typography
          className="articleCard__title"
          gutterBottom
          variant="h5"
          component="div"
        >
          {props.articleTitle}
        </Typography>
        <Typography
          className="articleCard__source"
          variant="body2"
          color="text.secondary"
        >
          {props.articleSourceTitle}
        </Typography>
        <Typography
          className="articleCard__shortContent"
          variant="body1"
          color="text.secondary"
        >
          {props.content}
        </Typography>
      </CardContent>
      <CardActions
        disableSpacing
        className={props.isAdmin ? 'd-none' : 'd-block'}
      >
        <IconButton aria-label="read later">
          <StarBorderIcon />
        </IconButton>
        <IconButton aria-label="status">
          <RadioButtonUncheckedIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
};

export default ContentCard;
