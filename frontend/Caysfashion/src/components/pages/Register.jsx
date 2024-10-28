import { LoginLayout } from "../layout/LoginLayout";
import { AuthInput } from "../fragments/AuthInput";
import { Button } from "../elements/Button";
import { AuthFormLayout } from "../layout/AuthFormLayout";
import { FooterLayout } from "../layout/FooterLayout";
import { Link } from "../elements/Link";
import { NavLayout } from "../layout/NavLayout";
import "../styles/Register.css";

export const RegisterPage = () => {
    return (
        <>
            <LoginLayout>
                <NavLayout className="register-nav">
                    <Link
                    href="/login"
                    id="register">Login</Link>
                    <Link
                    href=""
                    id="explore">Explore About</Link>
                </NavLayout>
                <AuthFormLayout
                className="register-form"
                titleForm="Register"
                guide="Already Have Account?"
                guideType="Login"
                guideLink="/login">
                    <div className="register-form-input">
                        <AuthInput
                        className="register-username"
                        id="username"
                        name="username"
                        placeholder="Username"
                        type="text">Username</AuthInput>
                        <AuthInput
                        className="register-email"
                        id="email"
                        name="email"
                        placeholder="Email"
                        type="email">Email</AuthInput>
                        <AuthInput
                        className="register-password"
                        id="password"
                        name="password"
                        placeholder="Password"
                        type="password">Password</AuthInput>
                        <AuthInput
                        className="register-password2"
                        id="password2"
                        name="password2"
                        placeholder="Confirm Password"
                        type="password">Confirm Password</AuthInput>
                    </div>
                    <div className="register-form-button">
                        <Button
                        className=""
                        type="">Register</Button>
                    </div>
                </AuthFormLayout>
                <FooterLayout className="register-footer"></FooterLayout>
            </LoginLayout>
        </>
    )
}