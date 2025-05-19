using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

using System;
using System.Diagnostics;

using SpiritLabelLibrary;

namespace demo_win_gui;

/// <summary>
/// Interaction logic for MainWindow.xaml
/// </summary>
public partial class MainWindow : Window
{
    public MainWindow()
    {
        InitializeComponent();
        InitializePrinterListComboBox();
    }
    
    private void InitializePrinterListComboBox()
    {
        foreach (var prn in SpiritLabel.PrnLst()) {
            cmbOptions.Items.Add(prn.name);
        }
        cmbOptions.SelectedIndex = 0; // 设置默认选中项
    }
	
	// 事件处理程序
	private void BtnPrint_Click(object sender, RoutedEventArgs e)
	{
	    var printer=cmbOptions.Items[cmbOptions.SelectedIndex] as string;
		
		var vars = new Dictionary<string, object>
        {
            { "co_name", "打印精灵" },
            { "name", "标签打印机" }
        };

        try {
			var p = SpiritLabel.OpenPrinter(printer);
			p.size(500, 300);
			p.Print("acae8013-28db-4b77-a500-1a6052633a22", vars);
			p.Close();
			MessageBox.Show($"打印完成", "提示");
        }catch(SpiritException ex) {
			MessageBox.Show($"打印错误: {ex.Message}", "提示");
		}
	}
	
	private void BtnEdit_Click(object sender, RoutedEventArgs e)
	{
		var file="c:\\tmp\\xxx.psl";
		SpiritLabel.NewLabel(file, "名称", "说明", 800, 1200, 300);
		SpiritLabel.Design(file);
	}
}
