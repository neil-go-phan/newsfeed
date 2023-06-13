type ChangePasswordFormProperty = {
  password: string;
  password_confirmation?: string;
}

type LoginFormProperty = {
  username: string;
  password: string;
}

type RegisterFormProperty = {
  email: string;
  username: string;
  password: string;
  passwordConfirmation?: string;
}

type UrlFormProperty = {
  url: string;
};

type Article = {
  title: string;
  description: string;
  link: string;
  published: string;
  authors:string;
}

type Articles = Array<Article>

type ArticlesSource = {
  title: string;
  description: string;
  link: string;
  feed_link: string;
  image: string;
}

type Crawler = {
  source_link: string;
  feed_link: string;
  crawl_type: string;
  article_div: string;
  article_title: string;
  article_description: string;
  article_link: string;
  article_published: string;
  article_authors: string;
  schedule: string;
}