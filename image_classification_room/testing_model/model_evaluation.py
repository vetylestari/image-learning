from tensorflow.keras.models import load_model

# Load the trained model
model = load_model('trained_model.h5')

# Evaluate the model on the test data
loss, accuracy = model.evaluate(
    test_generator,
    steps=len(test_generator)
)

# Print the test results
print(f'Test loss: {loss:.2f}')
print(f'Test accuracy: {accuracy * 100:.2f}%')