import protocol SwiftUI.App
import protocol SwiftUI.Scene

import struct FeatureZelda.IgnorableVoidApp

@main
struct App {
    // MARK: - Private members

    private let app = IgnorableVoidApp()
}

// MARK: - App

extension App: SwiftUI.App {
    var body: some Scene {
        app.body
    }
}
