export const OtpService = async (otpCode) => {
    try {
        const response = await fetch("http://localhost:8080/caysfashion/otp-verification", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ otp: otpCode }),
        });
        if (!response.ok) {
            const errorDetails = await response.json(); 
            throw new Error(errorDetails.error || "OTP verification failed");
        }

        return response.json(); 

    } catch (error) {
        console.error("Error posting OTP:", error);
        throw error; 
    }
};
