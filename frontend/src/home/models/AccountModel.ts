import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import authProvider from "../../dataprovider/authProvider";

async function checkAvailability({ username = '', email = '' })
{
    let result: boolean;

    try
    {
        let dataProvider = DataProviderFactory.getDataProvider()

        result = await dataProvider.create('register/availability', {
            data: {
                username: username,
                email: email,
            }
        })
        .then((resp:any) => {
            console.log(resp)
            return !resp?.data?.exist
        }, error => {
            NotificationManager.error(error.message, `Cek Ketersediaan Akun: Error Koneksi ${error.name}`);
            return true
        })


        if(!result)
        {
            NotificationManager.warning(
                `Akun "${username} ${email}" telah dipakai oleh orang lain atau telah terdaftar sebelumnya!`,
                'Akun tidak tersedia');
        }

    }
    catch (e)
    {
        result = true
    }

    return result
}

async function registerNewAccount(formData: any) {
    let dataProvider = DataProviderFactory.getDataProvider()
    try {
        return await dataProvider.create("temp_users", { data: formData }).then(_ => {
            NotificationManager.success(
                'Pendaftaran Sukses, Mohon Tunggu Konfirmasi Admin!', 'Pendaftaran Sukses')
            return true
        }, error => {
            if(error.message.trim().toLowerCase() === 'conflict') {
                NotificationManager.error(`Data telah terdaftar sebelumnya! ${error.message}`, 'Pendaftaran Gagal')
            } else {
                NotificationManager.error(error.message, `Pendaftaran Gagal: ${error.name}`)
            }
            return false
        })
    } catch (e) {
        NotificationManager.error(e.toString(), 'Kesalahan Teknis!');
        return false
    }
}

async function login({username, password}) {
    return await authProvider.login({username, password})
        .then(() => {
            NotificationManager.success(
                'Login Sukses, Mohon Tunggu Sebentar!', 'Login Sukses')
            return true
        })
        .catch(() => {
            NotificationManager.error(
                'Username/Password salah atau kendala koneksi jaringan!', 'Login Gagal')
            return false
        })
}

export {
    checkAvailability,
    registerNewAccount,
    login
}
