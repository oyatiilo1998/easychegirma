import AuthLayout from "../layouts/auth/Auth.layout";
import LoginPage from "../views/Login";

const authRoutes = [
  {
    component: LoginPage,
    path: "/login",
    exact: true,
    title: "Login",
    layout: AuthLayout
  },
  // {
  //   component: RegistrationPage,
  //   path: "/registration",
  //   exact: true,
  //   title: "Registration",
  //   layout: AuthLayout
  // },
  // {
  //   component: LandingPage,
  //   path: "/",
  //   exact: true,
  //   title: "LandingPage",
  //   layout: LandingLayout
  // },
];

export default authRoutes;
