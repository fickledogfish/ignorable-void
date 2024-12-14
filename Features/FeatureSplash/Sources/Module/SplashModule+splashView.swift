import protocol SwiftUI.View

extension SplashModule {
    public func splashView() -> some View {
        SplashView(
            SplashPresentation.self,
            presenter: SplashPresentation()
        )
    }
}
