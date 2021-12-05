import CategoriesList from "../views/Categories";
import CategoryCreate from "../views/Categories/Create";
import ProductsList from "../views/Products";
import ProductCreate from "../views/Products/Create";
import UsersList from "../views/Users";
import UserCreate from "../views/Users/Create";

const mainRoutes = [
  {
    component: CategoriesList,
    path: "/categories",
    exact: true,
    title: "ChaptersList"
  },
  {
    component: CategoryCreate,
    path: "/categories/create",
    exact: true,
    title: "CategoriesCreate"
  },
  {
    component: CategoryCreate,
    path: "/categories/edit/:id",
    exact: true,
    title: "CategoriesEdit"
  },
  {
    component: ProductsList,
    path: "/products",
    exact: true,
    title: "ProductsList"
  },
  {
    component: ProductCreate,
    path: "/products/create",
    exact: true,
    title: "ProductCreate"
  },
  {
    component: ProductCreate,
    path: "/products/edit/:id",
    exact: true,
    title: "ProductEdit"
  },
  {
    component: UsersList,
    path: "/users",
    exact: true,
    title: "UsersList"
  },
  {
    component: UserCreate,
    path: "/users/create",
    exact: true,
    title: "UserCreate"
  },
  {
    component: UserCreate,
    path: "/users/edit/:id",
    exact: true,
    title: "UserEdit"
  },
].map((route) => ({
  ...route,
  path: `/main${route.path}`,
  id: Math.random() + new Date().getTime(),
}));

export default mainRoutes