import { Model } from "react-model";

import AboutModel from "./AboutModel";
import ArticleTopicModel from "./ArticleTopicModel";

const models = { AboutModel, ArticleTopicModel }

export const { getInitialState, useStore, getState, actions, subscribe, unsubscribe } = Model(models)
