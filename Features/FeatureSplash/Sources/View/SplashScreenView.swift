import Combine
import SwiftUI

import class DesignTheme.DesignTheme
import class DesignTheme.VariableDesignTheme

@MainActor
struct SplashScreenView<Presenter: SplashScreenPresenter> {
    // MARK: - Private members

    @EnvironmentObject
    private var theme: DesignTheme

    @StateObject
    private var presenter: Presenter

    @State
    private var scale: CGFloat = 200

    // MARK: - Initializers

    init(
        _: Presenter.Type = Presenter.self,
        presenter: Presenter
    ) {
        self._presenter = StateObject(wrappedValue: presenter)
    }
}

// MARK: - View

extension SplashScreenView: View {
    var body: some View {
        VStack(spacing: theme.cgFloat(spacing: .large)) {
            Text("Some App")
                .font(.title)
                .background {
                    Color.blue
                        .clipShape(Circle())
                        .scaleEffect(scale)
                        .ignoresSafeArea()
                }

            HStack(alignment: .bottom, spacing: .zero) {
                Text("(not) ")
                    .font(.footnote)

                Text("Made by Very Smart People©")
                    .font(.callout)
            }
        }
        .padding()
    }
}

// MARK: - Previews

private enum Previews {
    final class Presenter: SplashScreenPresenter {
    }
}

#Preview {
    SplashScreenView(
        Previews.Presenter.self,
        presenter: Previews.Presenter()
    )
    .designTheme(VariableDesignTheme())
}
