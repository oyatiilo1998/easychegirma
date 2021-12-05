import CategoryPage from "../pages/category";
import MainPage from "../pages/main";


export const routesList = [
  {
    key: "main",
    path: "/",
    component: MainPage,
    exact: true
  },
  {
    key: "categoryPage",
    path: "/category/:id",
    component: CategoryPage,
    exact: true
  }
]