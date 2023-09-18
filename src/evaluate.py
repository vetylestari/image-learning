import tensorflow as tf
from model import create_model

def evaluate_model(test_dataset):
    input_shape = (your_input_shape)  # Gantilah dengan bentuk input yang sesuai
    num_classes = (your_num_classes)  # Gantilah dengan jumlah kelas yang sesuai

    model = create_model(input_shape, num_classes)

    # Muat model yang telah disimpan
    model = tf.keras.models.load_model('models/saved_model/my_model')

    # Evaluasi model pada dataset pengujian
    loss, accuracy = model.evaluate(test_dataset)
    
    return loss, accuracy