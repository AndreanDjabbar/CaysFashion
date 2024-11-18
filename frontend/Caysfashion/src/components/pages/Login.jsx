import { AuthLayout } from "../layouts/AuthLayout";
import { AuthInput } from "../fragments/AuthInput";
import { Button } from "../elements/Button";
import { AuthFormLayout } from "../layouts/AuthFormLayout";
import { FooterLayout } from "../layouts/FooterLayout";
import { Link } from "../elements/Link";
import { NavLayout } from "../layouts/NavLayout";
import "../../styles/Login.css";

export const LoginPage = () => {
    return (
        <>
            <AuthLayout>
                <NavLayout className="login-nav">
                    <Link
                    href="/register"
                    id="register">Register</Link>
                    <Link
                    href=""
                    id="explore">Explore About</Link>
                </NavLayout>
                <AuthFormLayout
                className="login-form"
                titleForm="Login"
                guide="Doesnt Have Account?"
                guideType="Register"
                guideLink="/register">
                    <div className="login-form-input">
                        <AuthInput
                        className="login-username"
                        id="username"
                        name="username"
                        placeholder="Username"
                        type="text">Username</AuthInput>
                        <AuthInput
                        className="login-password"
                        id="password"
                        name="password"
                        placeholder="Password"
                        type="password">Password</AuthInput>
                    </div>
                    <div className="login-form-button">
                        <Button
                        className=""
                        type="">Login</Button>
                    </div>
                </AuthFormLayout>
                <FooterLayout className="login-footer"></FooterLayout>
            </AuthLayout>
        </>
    )
}