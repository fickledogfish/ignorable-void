import SwiftUI

public protocol SplashRouter {
    associatedtype SplashScreenView: View

    func splashScreen() -> SplashScreenView
}

final class GeneralRouter {
    let splash: any SplashRouter

    init(splash: any SplashRouter) {
        self.splash = splash
    }

    func route() {
        splash.splashScreen()
    }
}

public final class DefaultSplashRouter {
    // MARK: - Initializers

    public init() {
    }
}

extension DefaultSplashRouter: SplashRouter {
    public func splashScreen() -> some View {
        EmptyView()
    }
}
