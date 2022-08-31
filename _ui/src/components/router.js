import { BrowserRouter, Routes, Route } from "react-router-dom";
import Auth from "./auth";
import Overview from "./overview";
import Menu from "./menu";
import NotFound from "./not_found";
import Targets from "./targets";
import AddTarget from "./add_target";

export default function Router() {

    return <BrowserRouter>
        <Routes>
            <Route
                path="/login"
                element={<Auth />}
            />
            <Route
                path="/"
                element={
                    <Menu viewElement={<Overview />} />
                }
            />
            <Route
                path="/targets"
                element={
                    <Menu viewElement={<Targets />} />
                }
            />
            <Route
                path="/targets/new"
                element={
                    <Menu viewElement={<AddTarget />} />
                }
            />
            <Route exact path="*" element={<NotFound />}>
            </Route>
        </Routes>
    </BrowserRouter>
}