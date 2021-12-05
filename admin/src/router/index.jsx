import { Switch, Route, Redirect, useLocation } from "react-router-dom"
import authRoutes from "./authRoutes"
import { useSelector } from "react-redux"
import { animated, useTransition } from "react-spring"
import AuthLayout from "../layouts/auth/Auth.layout"
import MainLayout from "../layouts/main/Main.layout"
import mainRoutes from "./mainRoutes"

const noAccessComponent = () => <>no access</>

const layouts = [{ component: MainLayout, path: "/main", routes: mainRoutes }]

const AppRouter = () => {
  const isAuth = useSelector((state) => state.auth.isAuth)
  const location = useLocation()
  const transitions = useTransition(location, {
    from: { opacity: 0 },
    enter: { opacity: 1 },
    leave: { opacity: 0 },
  })

  if (!isAuth)
    return (
      <Switch>
        {authRoutes.map((route, index) => (
            <Route
              path={route.path}
              exact={route.exact}
              key={index}
              render={routeProps => (
                // <route.layout>
                  <route.component {...routeProps}/>
                // </route.layout>
              )}
            />
        ))}
        <Redirect to="/login" />
      </Switch>
    )
  return (
    <Switch>
      {layouts.map((layout, index) => (
        <Route
          key={index}
          path={layout.path}
          render={(routeProps) => (
            <layout.component>
              {transitions((props, item) => (
                <animated.div style={props}>
                  <div style={{ position: "absolute", width: "100%" }}>
                    <Switch location={item}>
                      {layout.routes.map((route) => (
                        <Route
                          key={route.id}
                          path={route.path}
                          component={route.component}
                          exact
                        />
                      ))}
                      <Redirect from="*" to="/main/categories" />
                    </Switch>
                  </div>
                </animated.div>
              ))}
            </layout.component>
          )}
        ></Route>
      ))}
      <Redirect from="/" to="/main/categories" />
      <Redirect from="*" to="/main/categories" />
    </Switch>
  )
}

export default AppRouter
