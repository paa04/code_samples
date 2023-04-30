# Python program to create
# a file explorer in Tkinter

# import all components
# from the tkinter library
from tkinter import *
from tkinter import messagebox as mb

import PIL.Image

import back

from PIL import *

# import filedialog module
from tkinter import filedialog

# Function for opening the
# file explorer window

path = ""
path_s = ""


def info_mes():
    mb.showinfo('Инфа', 'Автор - Парфенов Арсений ИУ-24Б\n'
                        'Программа кодирует строчку в изображение,\n'
                        'также можно применить раскодирование по картинке')


def over_pic():
    mb.showinfo('Ошибка', 'Слишком маленькая картинка для этой строки')


def empt_path():
    mb.showinfo('Ошибка', 'Вы не выбрали файл!!!')


def save_ask(im):
    answer = mb.askyesno(
        title="Сохранение",
        message="Сохранить картинку?")
    if answer:
        saveChose()
        print(path_s)
        im.save(path_s)


def code_str():
    if (path != ""):
        image = PIL.Image.open(path)
        str = entry_str.get()
        im = back.str_code(str, image)
        if im != None:
            save_ask(im)
        else:
            over_pic()
    else:
        empt_path()


def decode():
    if path != "":
        image = PIL.Image.open(path)
        str = back.decode_pic(image)
        if str != None:
            str_dec.configure(text="Your str: " + str)
        else:
            str_dec.configure(text="No coded str in pic")
    else:
        empt_path()


def browseFiles():
    global path
    path = filedialog.askopenfilename(initialdir="/",
                                      title="Select a File",
                                      filetypes=(("Pic files", "*.bmp"),))
    # Change label contents
    # label_file_explorer.configure(text="File Opened: " + filename)


def saveChose():
    global path_s
    path_s = filedialog.asksaveasfilename(initialdir="/",
                                          title="Select a File",
                                          filetypes=(("Pic files", "*.bmp"),)) + '.bmp'


# Create the root window
window = Tk()

# Set window title
window.title('Pic coder')

# Set window size
window.geometry("900x500")

# Set window background color
window.config(background="white")

label_func = Label(window, text='Введите строчку: ')
label_func.grid(row=0, column=0, sticky='e')
entry_str = Entry(window, width=100)
entry_str.grid(row=0, column=1)

# Create a File Explorer label
str_dec = Label(window,
                text="",
                width=100, height=2,
                )

button_explore = Button(window, width=30,
                        text="Выбрать файл",
                        command=browseFiles)

button_code = Button(window, width=30,
                     text="Зашифровать",
                     command=code_str)

button_decode = Button(window, width=30, text="Расшифровать", command=decode)

button_exit = Button(window, width=30,
                     text="Закрыть",
                     command=exit)

# Grid method is chosen for placing
# the widgets at respective positions
# in a table like structure by
# specifying rows and columns
str_dec.grid(column=1, row=1)

button_explore.grid(column=1, row=2)

button_code.grid(column=1, row=3)
button_decode.grid(column=1, row=4)
button_exit.grid(column=1, row=5)

mainmenu = Menu(window)
window.config(menu=mainmenu)
mainmenu.add_command(label="Информация о программе", command=info_mes)
# Let the window wait for any events
window.mainloop()
