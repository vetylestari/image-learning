import tensorflow as tf
from tensorflow.keras import losses, optimizers
from data_preprocessing import preprocess_dataset
from model import create_model

def train_model(train_dataset, val_dataset, num_epochs, learning_rate):
    input_shape = (your_input_shape)  # Gantilah dengan bentuk input yang sesuai
    num_classes = (your_num_classes)  # Gantilah dengan jumlah kelas yang sesuai

    train_dataset = preprocess_dataset(train_dataset)
    val_dataset = preprocess_dataset(val_dataset)

    model = create_model(input_shape, num_classes)

    model.compile(optimizer=optimizers.Adam(learning_rate),
                  loss=losses.SparseCategoricalCrossentropy(),
                  metrics=['accuracy'])

    history = model.fit(train_dataset,
                        validation_data=val_dataset,
                        epochs=num_epochs)
    
    # Simpan model ke direktori models/saved_model/
    model.save('models/saved_model/my_model')

    return history