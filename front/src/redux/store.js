import { configureStore } from "@reduxjs/toolkit";
import { combineReducers } from "redux";
import thunk from "redux-thunk";
import { mainReducer } from "./slices/main.slice";


const rootReducer = combineReducers({
  main: mainReducer
})

const store = configureStore({
  reducer: rootReducer,
  middleware: [thunk]
})

export default store