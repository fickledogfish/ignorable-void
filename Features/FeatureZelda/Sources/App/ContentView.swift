import SwiftUI

import DesignTheme

import class FeatureSplash.SplashModule

@MainActor
struct ContentView {
    let splash = SplashModule()
}

// MARK: - View

extension ContentView: View {
    var body: some View {
        NavigationStack {
            splash.splashView()
        }
    }
}

// MARK: - Previews

#Preview {
    ContentView()
        .designTheme(VariableDesignTheme())
}
