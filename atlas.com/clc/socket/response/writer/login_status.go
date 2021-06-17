package writer

import (
	"atlas-clc/socket/response"
	"github.com/sirupsen/logrus"
)

const OpCodeLoginStatus uint16 = 0x00

func WriteAuthSuccess(l logrus.FieldLogger) func(accountId uint32, name string, gender byte, pic string) []byte {
	return func(accountId uint32, name string, gender byte, pic string) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeLoginStatus)
		w.WriteInt(0)
		w.WriteShort(0)
		w.WriteInt(accountId)
		w.WriteByte(gender)
		//boolean canFly = false;// Server.getInstance().canFly(client.getAccID());
		//writer.writeBool((YamlConfig.config.server.USE_ENFORCE_ADMIN_ACCOUNT || canFly) && client.getGMLevel() > 1);    // GM
		w.WriteBool(false)
		//writer.write(((YamlConfig.config.server.USE_ENFORCE_ADMIN_ACCOUNT || canFly) && client.getGMLevel() > 1) ? 0x80 : 0);  //
		// Admin Byte. 0x80,0x40,0x20.. Rubbish.
		w.WriteByte(0)
		// country code
		w.WriteByte(0)
		w.WriteAsciiString(name)
		w.WriteByte(0)
		// quiet ban
		w.WriteByte(0)
		// quiet ban timestamp
		w.WriteLong(0)
		// creation timestamp
		w.WriteLong(0)
		// 1 : remove the select the world you want to play in
		w.WriteInt(1)
		// 0 = Pin-System Enabled, 1 = Disabled
		w.WriteByte(1)
		// 0 = Register PIC, 1 = Ask for PIC, 2 = Disabled
		w.WriteByte(2)
		return w.Bytes()
	}
}

func WriteTemporaryBan(l logrus.FieldLogger) func(until uint64, reason byte) []byte {
	return func(until uint64, reason byte) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeLoginStatus)
		w.WriteByte(2)
		w.WriteByte(0)
		w.WriteInt(0)
		w.WriteByte(reason)
		// Temp ban date is handled as a 64-bit long, number of 100NS intervals since 1/1/1601.
		w.WriteLong(until)
		return w.Bytes()
	}
}

func WritePermanentBan(l logrus.FieldLogger) []byte {
	w := response.NewWriter(l)
	w.WriteShort(OpCodeLoginStatus)
	w.WriteByte(2)
	w.WriteByte(0)
	w.WriteInt(0)
	w.WriteByte(0)
	w.WriteLong(0)
	return w.Bytes()
}

func WriteLoginFailed(l logrus.FieldLogger) func(reason byte) []byte {
	return func(reason byte) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeLoginStatus)
		w.WriteByte(reason)
		w.WriteByte(0)
		w.WriteInt(0)
		return w.Bytes()
	}
}
